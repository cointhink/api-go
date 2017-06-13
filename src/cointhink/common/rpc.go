package common

import "encoding/json"
import "strings"
import "reflect"
import "log"

import "cointhink/model"

import "github.com/golang/protobuf/jsonpb"
import gproto "github.com/golang/protobuf/proto"
import "github.com/gorilla/websocket"

//import "github.com/satori/go.uuid"
import "github.com/google/uuid"

// rpc
var RPCq chan RpcMsg
var OUTq chan RpcOut

type RpcMsg struct {
	socket    *websocket.Conn
	accountId string
	payload   []byte
}

type RpcResponse struct {
	msg *gproto.Message
	id  string
}

type RpcOut struct {
	socket    *websocket.Conn
	accountId string
	response  *RpcResponse
}

func Rpc(msg *RpcMsg) {
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

	log.Printf("response: %p %d msg", msg.socket, len(responses))
	for _, response := range responses {
		OUTq <- RpcOut{socket: msg.socket,
			response: &RpcResponse{msg: &response, id: dat["id"].(string)}}
	}
}

func Respond(out *RpcOut) {
	response_class := reflect.TypeOf(*out.response.msg).String()
	method := strings.Split(response_class, ".")[1]
	marsh := jsonpb.Marshaler{}
	objJson, err := marsh.MarshalToString(*out.response.msg)
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
	resp := map[string]interface{}{"id": out.response.id,
		"method": method,
		"object": jsonified}
	json, err := json.Marshal(resp)
	if err != nil {
		log.Println("tojson:", err)
		return
	}
	log.Printf("ws_send: %p %s", out.socket, json)
	err = out.socket.WriteMessage(websocket.TextMessage, json)
	if err != nil {
		log.Println("ws_send err:", err)
		return
	}
}

func RpcId() string {
	uuid, _ := uuid.NewRandom()
	uuidStr := uuid.String()
	return uuidStr[16:35]
}
