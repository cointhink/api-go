package actions

import (
	"log"

	"db"
	"proto"

	"github.com/golang/protobuf/jsonpb"
)

func DoSignupform(form proto.SignupForm, json string) proto.SignupFormResponse {
	err := jsonpb.UnmarshalString(json, &form)
	if err != nil {
		log.Print("unmarshaling error: ", err)
		return proto.SignupFormResponse{Ok: true}
	}

	rows, err := db.D.Handle.Query(
		"insert into accounts (id, email, username, fullname) values ($1, $2, $3, $4)",
		db.NewId("accounts"),
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
