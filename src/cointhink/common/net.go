package common

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Httpclient struct {
	socket    *websocket.Conn
	accountId string
	out       []*RpcOut
}

var clients map[*websocket.Conn]Httpclient

func Httploop(listen_address string) {
	log.Printf("http listening %s", listen_address)
	http.HandleFunc("/", Upgrade)
	clients = map[*websocket.Conn]Httpclient{}
	http.ListenAndServe(listen_address, nil)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s webstocket started", r.Header.Get("Origin"))
	wsocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("websocket upgrade fail:", err)
		return
	}
	_client := Httpclient{socket: wsocket, out: []*RpcOut{}}
	log.Printf("wsocket open %p", wsocket)
	clients[wsocket] = _client
	for {
		_, payload, err := wsocket.ReadMessage()
		if err != nil {
			log.Println("ws_recv err:", err)
			break
		}
		RPCq <- RpcMsg{socket: _client.socket, payload: payload}
	}
	log.Printf("wsocket closing %p", wsocket)
	wsocket.Close()
	delete(clients, wsocket)
}

func AccountToSocket(accountId string) *websocket.Conn {
	for _, c := range clients {
		if c.accountId == accountId {
			return c.socket
		}
	}
	return nil
}
