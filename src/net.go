package main

import (
    "net/http"
    "encoding/json"
    "log"

    "pb"

    "github.com/golang/protobuf/jsonpb"
    "github.com/gorilla/websocket"
)

type WsResponse struct {
    id    string
}

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) {
    log.Print("upgrading", r.Header.Get("Origin"))
    c, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade:", err)
        return
    }
    defer c.Close()
    for {
        mt, message, err := c.ReadMessage()
        if err != nil {
            log.Println("read:", err)
            break
        }
        log.Printf("recv: %s", message)

        var dat map[string]interface{}

        json.Unmarshal(message, &dat)
        log.Printf("%+v", dat["object"])
        object_json, _ := json.Marshal(dat["object"])
        newTest := &signup_form.SignupForm{}
        err = jsonpb.UnmarshalString(string(object_json), newTest)
        if err != nil {
            log.Print("unmarshaling error: ", err)
        }
        log.Printf("newTest: %+v", newTest)

        resp := map[string]interface{}{"id": "Ab12", "object": newTest}
        json, _ := json.Marshal(resp)
        log.Printf("returning: %s", json)
        err = c.WriteMessage(mt, json)
        if err != nil {
            log.Println("write:", err)
            break
        }
    }
}