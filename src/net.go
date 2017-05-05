package main

import (
    "net/http"
    "encoding/json"
    "log"

    "github.com/gorilla/websocket"
)

type WsResponse struct {
    id    string
    object interface{}
}

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) {
    log.Print("upgrading", r.Header.Get("Origin"))
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
        log.Printf("recv: %s", payload)

        var dat map[string]interface{}

        json.Unmarshal(payload, &dat)
        log.Printf("%+v", dat["object"])

        response_class, response_object := Dispatch(dat["method"].(string), dat["object"])

        //resp := WsResponse{"Ab12", newTest} ??
        resp := map[string]interface{}{"id": dat["id"], "method": response_class,
                                       "object": response_object }
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