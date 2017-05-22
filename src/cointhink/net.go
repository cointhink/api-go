package cointhink

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gorilla/websocket"
)

type WsResponse struct {
	id     string
	object interface{}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) {
	log.Printf("http websocket upgrade from %s", r.Header.Get("Origin"))
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("websocket upgrade fail:", err)
		return
	}
	defer c.Close()
	for {
		_, payload, err := c.ReadMessage()
		if err != nil {
			log.Println("ws_recv err:", err)
			break
		}

		var dat map[string]interface{}
		json.Unmarshal(payload, &dat)

		responses := Dispatch(dat["method"].(string), dat["object"])
		for _, response := range responses {
			respond(c, response, dat["id"].(string))
		}

	}
}

func respond(client *websocket.Conn, response interface{}, id string) {
	response_class := reflect.TypeOf(response).String()
	method := strings.Split(response_class, ".")[1]
	resp := map[string]interface{}{"id": id,
		"method": method,
		"object": response}
	log.Printf("resp wtf %#v", response)
	json, err := json.Marshal(resp)
	if err != nil {
		log.Println("tojson:", err)
		return
	}
	log.Printf("ws_send: %s", json)
	err = client.WriteMessage(websocket.TextMessage, json)
	if err != nil {
		log.Println("ws_send:", err)
		return
	}
}
