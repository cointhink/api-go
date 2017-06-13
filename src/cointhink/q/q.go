package q

import gproto "github.com/golang/protobuf/proto"
import "github.com/gorilla/websocket"

var OUTq chan RpcOut

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
	Socket    *websocket.Conn
	AccountId string
	Response  *RpcResponse
}
