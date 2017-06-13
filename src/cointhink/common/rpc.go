package common

import "encoding/json"
import "strings"
import "reflect"
import "log"

import "cointhink/model"
import "cointhink/q"
import "cointhink/httpclients"

import "github.com/golang/protobuf/jsonpb"
import gproto "github.com/golang/protobuf/proto"
import "github.com/gorilla/websocket"

// rpc
var RPCq chan q.RpcMsg

func Rpc(msg *q.RpcMsg) {
	var dat map[string]interface{}
	json.Unmarshal(msg.Payload, &dat)
	method := dat["method"].(string)
	objectBytes, _ := json.Marshal(dat["object"])
	objectJson := string(objectBytes)

	var responses []gproto.Message
	responses = DispatchPublic(method, objectJson)
	if responses == nil {
		if dat["token"] != nil {
			token := dat["token"].(string)
			accountId, err := model.TokenFindAccountId(token)
			if err != nil {
				log.Printf("msg token %s BAD", token)
				return
			}
			httpclient := httpclients.Clients[msg.Socket]
			httpclient.AccountId = accountId
			httpclients.Clients[msg.Socket] = httpclient
			responses = DispatchAuth(method, objectJson, accountId)
		}
	}

	log.Printf("response: %p %d msg", msg.Socket, len(responses))
	for _, response := range responses {
		q.OUTq <- q.RpcOut{Socket: msg.Socket,
			Response: &q.RpcResponse{Msg: response, Id: dat["id"].(string)}}
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
	log.Printf("ws_send: %p %s", out.Socket, json)
	err = out.Socket.WriteMessage(websocket.TextMessage, json)
	if err != nil {
		log.Println("ws_send err:", err)
		return
	}
}
