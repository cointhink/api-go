package common

import "encoding/json"
import "strings"
import "reflect"
import "log"

import "cointhink/model"

import "github.com/golang/protobuf/jsonpb"
import gproto "github.com/golang/protobuf/proto"
import "github.com/gorilla/websocket"

// rpc
var RPCq chan RpcMsg
var OUTq chan Httpclient

type RpcMsg struct {
	client  Httpclient
	payload []byte
}

type RpcOut struct {
	client Httpclient
	msg    gproto.Message
	id     string
}

func Rpc(msg RpcMsg) {
	var dat map[string]interface{}
	json.Unmarshal(msg.payload, &dat)
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
			responses = DispatchAuth(method, objectJson, accountId)
		}
	}

	log.Printf("response: %d msg", len(responses))
	for _, response := range responses {
		msg.client.out = append(msg.client.out, RpcOut{client: msg.client,
			msg: response,
			id:  dat["id"].(string)})
	}
	OUTq <- msg.client // signal to flush this client's out queue
}

func Respond(socket *websocket.Conn, response gproto.Message, id string) {
	response_class := reflect.TypeOf(response).String()
	method := strings.Split(response_class, ".")[1]
	marsh := jsonpb.Marshaler{}
	objJson, err := marsh.MarshalToString(response)
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
	resp := map[string]interface{}{"id": id,
		"method": method,
		"object": jsonified}
	json, err := json.Marshal(resp)
	if err != nil {
		log.Println("tojson:", err)
		return
	}
	log.Printf("ws_send: %s", json)
	err = socket.WriteMessage(websocket.TextMessage, json)
	if err != nil {
		log.Println("ws_send:", err)
		return
	}
}