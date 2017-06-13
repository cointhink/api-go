package httpclients

import "github.com/gorilla/websocket"
import "cointhink/q"

type Httpclient struct {
	Socket    *websocket.Conn
	AccountId string
	Out       []*q.RpcOut
}

var Clients map[*websocket.Conn]Httpclient

func AccountIdToSocket(accountId string) *websocket.Conn {
	for _, c := range Clients {
		if c.AccountId == accountId {
			return c.Socket
		}
	}
	return nil
}
