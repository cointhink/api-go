package httpclients

import "github.com/gorilla/websocket"
import "cointhink/q"

var (
	Pinglist []string
)

type Httpclient struct {
	Socket    *websocket.Conn
	AccountId string
	AlgorunId string
	Out       []*q.RpcOut
}

var Clients map[*websocket.Conn]Httpclient

func AccountIdToWebSockets(accountId string) []*websocket.Conn {
	var list []*websocket.Conn
	for _, c := range Clients {
		if c.AccountId == accountId && c.AlgorunId == "" {
			list = append(list, c.Socket)
		}
	}
	return list
}
