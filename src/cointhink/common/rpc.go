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
		if responses == nil {
			token_, err := token.FindByToken(call.Token)
			if err != nil {
				log.Printf("common.Rpc token %s BAD %+v", call.Token, err)
				return
			}
			httpclient := httpclients.Clients[msg.Socket]
			httpclient.AccountId = token_.AccountId
			httpclient.AlgorunId = token_.AlgorunId
			httpclients.Clients[msg.Socket] = httpclient
			responses = DispatchAuth(call.Method, call.Object, token_.AccountId)
		}
	}
	log.Printf("rpc response: %p/%s %d msg", msg.Socket, msg.AccountId, len(responses))
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
	executors := LambdaExecutors()
	log.Printf("lambdaall has %d executors", len(executors))

	for _, _algorun := range algorun.Lambdable() {
		token, err := token.FindByAccountId(_algorun.AccountId, _algorun.Id)
		if err != nil {
		} else {
			for _, e := range executors {
				marketPrice_object, err := ptypes.MarshalAny(marketPrice)
				if err != nil {
					log.Printf("lambdaall algorun %+v schedule %+v", _algorun.Id, _algorun.ScheduleId)
					lambda := &proto.Lambda{
						Token:   token.Token,
						Method:  protoClassName(marketPrice),
						Object:  marketPrice_object,
						StateIn: _algorun.State}
					q.OUTq <- q.RpcOut{Socket: e.Socket,
						Response: &q.RpcResponse{Msg: lambda, Id: id}}
				}
			}
		}
	}
}

func LambdaExecutors() []*httpclients.Httpclient {
	executors := []*httpclients.Httpclient{}
	for _, client := range httpclients.Clients {
		if len(client.AlgorunId) > 0 {
			log.Printf("lambdaexecutors algorun %s", client.AlgorunId)
			_algorun, err := algorun.Find(client.AlgorunId)
			if err != nil {
				log.Printf("lambdaexecutors algorun schedule %s", _algorun.ScheduleId)
				_schedule, err := schedule.Find(_algorun.ScheduleId)
				if err != nil {
					log.Printf("lambdaexecutors algorun schedule executor %s algorithm", _schedule.Executor, _schedule.AlgorithmId)
					if _schedule.Executor == proto.Schedule_lambda {
						if err != nil {
							executors = append(executors, &client)
						}
					}
				}
			}
		}
	}
	return executors
}

func protoClassName(proto gproto.Message) string {
	response_class := reflect.TypeOf(proto).String()
	return strings.Split(response_class, ".")[1]
}

func Respond(out *q.RpcOut) {
	if out.Response == nil {
		if err := out.Socket.WriteMessage(websocket.PingMessage, []byte("cointhink")); err != nil {
			log.Printf("Ping send err %+v!\n", err)
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
