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
			ret = doSignupform(*obj)
		}
	}
	return ret
}

func doSignupform(form proto.SignupForm) proto.SignupFormResponse {
	rows, err := db.handle.Query(
		"insert into accounts (id, email, username, fullname) values ($1, $2, $3, $4)",
		NewId("accounts"),
		form.Account.GetEmail(),
		form.Account.GetUsername(),
		form.Account.GetFullname())
	if err != nil {
		log.Printf("upsert %+v", err)
		return proto.SignupFormResponse{Ok: false}
	}
	cols, err := rows.Columns()
	if err == nil {
		log.Printf("rows %+v", cols)
	} else {
		log.Printf("err %s", err)
	}
	resp := proto.SignupFormResponse{Ok: true}
	return resp
}
