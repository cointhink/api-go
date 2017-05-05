package main

import (
    "log"
    "encoding/json"

    "proto"

    "github.com/golang/protobuf/jsonpb"
)

type empty struct {}

func Dispatch(class string, object interface{}) interface{} {
    log.Printf("dispatch %s %+v", class, object)
    object_json, _ := json.Marshal(object)
    var ret interface{}
    switch class {
    case "SignupForm":
        obj := &proto.SignupForm{}
        err := jsonpb.UnmarshalString(string(object_json), obj)
        if err != nil {
            log.Print("unmarshaling error: ", err)
        }
        ret = doSignupform(obj)
    }

    return ret
}

func doSignupform(form *proto.SignupForm) proto.SignupFormResponse {
    resp := proto.SignupFormResponse{Ok: true}
    return resp
}