package common

import "encoding/json"
import "strings"
import "reflect"
import "log"

import "cointhink/model"
import "cointhink/q"
import "cointhink/httpclients"
import "cointhink/proto"

import "github.com/golang/protobuf/jsonpb"
import gproto "github.com/golang/protobuf/proto"
import "github.com/gorilla/websocket"

// rpc
var RPCq chan q.RpcMsg

type call struct {
	Id     string      `json:"id"`
	Method string      `json:"method"`
	Object interface{} `json:"object"`
	Token  string      `json:"token"`
}

func Rpc(msg *q.RpcMsg) {
	var responses []gproto.Message

	call := proto.Rpc{}
	err := jsonpb.UnmarshalString(string(msg.Payload), &call)
	if err != nil {
		log.Printf("ws rpc parse err:%+v", err)
	} else {
		responses = DispatchPublic(call.Method, call.Object)
		if responses == nil {
			accountId, err := model.TokenFindAccountId(call.Token)
			if err != nil {
				log.Printf("msg token %s BAD", call.Token)
				return
			}
			httpclient := httpclients.Clients[msg.Socket]
			httpclient.AccountId = accountId
			httpclients.Clients[msg.Socket] = httpclient
			responses = DispatchAuth(call.Method, call.Object, accountId)
		}
	}
	log.Printf("rpc response: %p/%s %d msg", msg.Socket, msg.AccountId, len(responses))
	for _, response := range responses {
		q.OUTq <- q.RpcOut{Socket: msg.Socket,
			Response: &q.RpcResponse{Msg: response, Id: call.Id}}
	}
}

func RespondAll(msg gproto.Message) {
	id := "rall"
	for _, client := range httpclients.Clients {
		q.OUTq <- q.RpcOut{Socket: client.Socket,
			Response: &q.RpcResponse{Msg: msg, Id: id}}
	}
}

func Respond(out *q.RpcOut) {
	response_class := reflect.TypeOf(out.Response.Msg).String()
	method := strings.Split(response_class, ".")[1]
	marsh := jsonpb.Marshaler{}
	objJson, err := marsh.MarshalToString(out.Response.Msg)
	if err != nil {
		log.Println("objJson:", err)
		return
	}
	var jsonified interface{}
	err = json.Unmarshal([]byte(objJson), &jsonified)
	if err != nil {
		log.Printf("unmah: %s", err)
		return
	}
	resp := map[string]interface{}{"id": out.Response.Id,
		"method": method,
		"object": jsonified}
	json, err := json.Marshal(resp)
	if err != nil {
		log.Println("tojson:", err)
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
