package common

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Httpclient struct {
	socket *websocket.Conn
}

var clients []Httpclient

func Httploop(listen_address string) {
	log.Printf("http listening %s", listen_address)
	http.HandleFunc("/", Upgrade)
	http.ListenAndServe(listen_address, nil)
}

type WsResponse struct {
	id     string
	object interface{}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s webstocket started", r.Header.Get("Origin"))
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("websocket upgrade fail:", err)
		return
	}
	_client := Httpclient{socket: c}
	clients = append(clients, _client)
	for {
		_, payload, err := c.ReadMessage()
		if err != nil {
			log.Println("ws_recv err:", err)
			break
		}
		RPCq <- RpcMsg{client: _client, payload: payload}
	}
	c.Close()
}
