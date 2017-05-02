package main

import (
    "net/http"
    "log"
    "pb"
    "github.com/golang/protobuf/jsonpb"
    "github.com/gorilla/websocket"
)

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

        newTest := &signup_form.SignupForm{}
        err = jsonpb.UnmarshalString(string(message), newTest)
        if err != nil {
            log.Print("unmarshaling error: ", err)
        }
        log.Printf("newTest: %+v", newTest)

        err = c.WriteMessage(mt, message)
        if err != nil {
            log.Println("write:", err)
            break
        }
    }
}