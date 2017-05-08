package main

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
		mt, payload, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("ws: %s", payload)

		var dat map[string]interface{}
		json.Unmarshal(payload, &dat)

		response_object := Dispatch(dat["method"].(string), dat["object"])
		response_class := reflect.TypeOf(response_object).String()
		method := strings.Split(response_class, ".")[1]
		resp := map[string]interface{}{"id": dat["id"],
			"method": method,
			"object": response_object}
		json, err := json.Marshal(resp)
		if err != nil {
			log.Println("write:", err)
			break
		}
		log.Printf("returning: %s", json)
		err = c.WriteMessage(mt, json)
		if err != nil {
			log.Println("write:", err)
			break
		}

	}
}
