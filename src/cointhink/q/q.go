package q

import gproto "github.com/golang/protobuf/proto"
import "github.com/gorilla/websocket"
import "github.com/google/uuid"

var OUTq chan RpcOut
var LXDOPq chan AccountOperation

type RpcMsg struct {
	Socket    *websocket.Conn
	AccountId string
	Payload   []byte
}

type RpcResponse struct {
	Msg gproto.Message
	Id  string
}

type RpcOut struct {
	Socket   *websocket.Conn
	Response *RpcResponse
}

func RpcId() string {
	uuid, _ := uuid.NewRandom()
	uuidStr := uuid.String()
	return uuidStr[19:35]
}
