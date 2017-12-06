package common

import "encoding/json"
import "strings"
import "reflect"
import "log"

import "cointhink/model/token"
import "cointhink/q"
import "cointhink/httpclients"
import "cointhink/proto"
import "cointhink/model/algorun"
import "cointhink/model/schedule"

import "github.com/golang/protobuf/jsonpb"
import gproto "github.com/golang/protobuf/proto"
import "github.com/gorilla/websocket"
import "github.com/golang/protobuf/ptypes"

// rpc
var RPCq chan q.RpcMsg

func Rpc(msg *q.RpcMsg) {
	var responses []gproto.Message

	call := proto.Rpc{}
	err := jsonpb.UnmarshalString(string(msg.Payload), &call)
	if err != nil {
		log.Printf("ws rpc parse err:%+v", err)
	} else {
		responses = DispatchPublic(call.Method, call.Object)
		if responses != nil {
			log.Printf("%s -> %#v (public)", msg.Socket.RemoteAddr().String(), call.Method)
		} else {
			token_, err := token.FindByToken(call.Token)
			if err != nil {
				log.Printf("common.Rpc token %s BAD %+v", call.Token, err)
				return
			}
			httpclient := httpclients.Clients[msg.Socket]
			httpclient.AccountId = token_.AccountId
			httpclient.AlgorunId = token_.AlgorunId
			httpclients.Clients[msg.Socket] = httpclient
			log.Printf("%s -> %#v (auth) AccountId:%#v Algorun:%#v", msg.Socket.RemoteAddr().String(), call.Method, token_.AccountId, token_.AlgorunId)
			responses = DispatchAuth(call.Method, call.Object, token_)
		}
	}
	log.Printf("%s <- %#v response %d msg", msg.Socket.RemoteAddr().String(), call.Method, len(responses))
	for _, response := range responses {
		q.OUTq <- q.RpcOut{Socket: msg.Socket,
			Response: &q.RpcResponse{Msg: response, Id: call.Id}}
	}
}

func RespondAll(msg gproto.Message) {
	id := "respondall"
	log.Printf("RespondAll http client count %d", len(httpclients.Clients))
	for _, client := range httpclients.Clients {
		q.OUTq <- q.RpcOut{Socket: client.Socket,
			Response: &q.RpcResponse{Msg: msg, Id: id}}
	}
}

func LambdaAll(marketPrice *proto.MarketPrices) {
	id := "lambdaall"
	lambdables := algorun.Lambdable()
	log.Printf("**lambdaall working on %d lambdables", len(lambdables))
	for _, _algorun := range lambdables {
		_schedule, err := schedule.Find(_algorun.ScheduleId)
		if err != nil {
		} else {
			executor := LambdaExecutor(_schedule.AlgorithmId)
			if executor == nil {
				log.Printf("!!lambdaall NO executor found for algorithm %s", _schedule.AlgorithmId)
			} else {
				log.Printf("lambdaall has executor %s for algorithm %s", _schedule.Id, _schedule.AlgorithmId)
				marketPrice_object, err := ptypes.MarshalAny(marketPrice)
				if err != nil {
				} else {
					token, err := token.FindByAccountId(_algorun.AccountId, _algorun.Id)
					if err != nil {
						log.Printf("lambdaall no token accountId %+v algorunId %+v", _algorun.AccountId,
							_algorun.Id)
					} else {
						log.Printf("lambdaall token algorun %+v schedule %+v", _algorun.Id, _algorun.ScheduleId)
						lambda := &proto.Lambda{
							Token:   token.Token,
							Method:  protoClassName(marketPrice),
							Object:  marketPrice_object,
							StateIn: _algorun.State}
						q.OUTq <- q.RpcOut{Socket: executor.Socket,
							Response: &q.RpcResponse{Msg: lambda, Id: id}}
					}
				}
			}
		}
	}
}

func LambdaExecutor(_algorithmId string) *httpclients.Httpclient {
	log.Printf("**lambdaexecutor search %d connected clients for lambda master algo %s", len(httpclients.Clients), _algorithmId)
	for _, client := range httpclients.Clients {
		if len(client.AlgorunId) > 0 {
			_algorun, err := algorun.Find(client.AlgorunId)
			if err != nil {
				log.Printf("!!lambdaexecutor algorun %s", err)
			} else {
				_schedule, err := schedule.Find(_algorun.ScheduleId)
				if err != nil {
					log.Printf("!!lambdaexecutor schedule %s", err)
				} else {
					if _schedule.AlgorithmId == _algorithmId {
						if _schedule.Executor == proto.Schedule_lambda_master {
							return &client
						} else {
							log.Printf("lambdaexecutors not master %d", _schedule.Executor)
						}
					} else {
						log.Printf("lambdaexecutors wrong algorithm %s", _schedule.AlgorithmId)
					}
				}
			}
		}
	}
	return nil
}

func protoClassName(proto gproto.Message) string {
	response_class := reflect.TypeOf(proto).String()
	return strings.Split(response_class, ".")[1]
}

func Respond(out *q.RpcOut) {
	if out.Response == nil {
		if err := out.Socket.WriteMessage(websocket.PingMessage, []byte("cointhink")); err != nil {
			log.Printf("Ping send err %+v! dropping client\n", err)
			httpclients.Clients.Remove(out.Socket)
		}
	} else {
		method := protoClassName(out.Response.Msg)
		jsonified := protoAnon(out.Response.Msg)
		resp := map[string]interface{}{"id": out.Response.Id,
			"method": method,
			"object": jsonified}
		json, err := json.Marshal(resp)
		if err != nil {
			log.Println("tojson err:", err)
			return
		}
		if out.Socket == nil {
			log.Println("ws_send err, socket is nil. aborted.")
		} else {
			log.Printf("ws_send: %p %s", out.Socket, json)
			err = out.Socket.WriteMessage(websocket.TextMessage, json)
			if err != nil {
				log.Println("ws_send err:", err)
				return
			}
		}
	}
}

func protoAnon(msg gproto.Message) interface{} {
	marsh := jsonpb.Marshaler{}
	objJson, err := marsh.MarshalToString(msg)
	if err != nil {
		log.Println("objJson:", err)
		return nil
	}
	var jsonified interface{}
	err = json.Unmarshal([]byte(objJson), &jsonified)
	if err != nil {
		log.Printf("unmarshal err: %s", err)
		return nil
	}
	return jsonified
}
