package common

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strings"

	"cointhink/model"

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
		method := dat["method"].(string)

		var responses []interface{}
		responses = DispatchPublic(method, dat["object"])
		if responses == nil {
			if dat["token"] != nil {
				token := dat["token"].(string)
				accountId, err := model.TokenFindAccountId(token)
				if err != nil {
					log.Printf("msg token %s BAD", token)
					return
				}
				responses = DispatchAuth(method, dat["object"], accountId)
			}
		}

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