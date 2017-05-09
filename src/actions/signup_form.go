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

	rows, _ := db.D.Handle.Query(
		"select count(*) from accounts where email = $1",
		form.Account.GetEmail())

	if rows.Next() {
		var count int
		rows.Scan(&count)
		if count > 0 {
			log.Printf("email check %+v", count)
			return proto.SignupFormResponse{Ok: false, Reason: proto.SignupFormResponse_EMAIL_TAKEN}
		}
	}

	rows, _ = db.D.Handle.Query(
		"select count(*) from accounts where username is not null and username = $1",
		form.Account.GetUsername())
	if rows.Next() {
		var count int
		rows.Scan(&count)
		if count > 0 {
			log.Printf("username check %+v", count)
			return proto.SignupFormResponse{Ok: false, Reason: proto.SignupFormResponse_USERNAME_TAKEN}
		}
	}

	rows, err = db.D.Handle.Query(
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
