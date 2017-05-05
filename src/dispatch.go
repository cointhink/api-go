package main

import (
    "log"
    "encoding/json"

    "proto"

    "github.com/golang/protobuf/jsonpb"
)

func Dispatch(class string, object interface{}) (string, interface{}) {
    log.Print()
    object_json, _ := json.Marshal(object)
    newTest := &proto.SignupForm{}
    err := jsonpb.UnmarshalString(string(object_json), newTest)
    if err != nil {
        log.Print("unmarshaling error: ", err)
    }
    log.Printf("newTest: %+v", newTest)
    sfrm := proto.SignupFormResponse{Ok: true}

    return "SignupFormResponse", sfrm
}