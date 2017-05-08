package main

import (
	"encoding/json"
	"log"

	"proto"

	"github.com/golang/protobuf/jsonpb"
)

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
			ret = proto.SignupFormResponse{Ok: true}
		} else {
			ret = doSignupform(obj)
		}
	}
	return ret
}

func doSignupform(form *proto.SignupForm) proto.SignupFormResponse {
	err := db.upsert(*form)
	if err != nil {
		log.Printf("upsert %+v", err)
	}
	resp := proto.SignupFormResponse{Ok: true}
	return resp
}
