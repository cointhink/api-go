package actions

import (
	"log"
	"strings"

	"db"
	"proto"

	"github.com/golang/protobuf/jsonpb"
)

func DoSignupform(form proto.SignupForm, json string) []interface{} {
	err := jsonpb.UnmarshalString(json, &form)
	if err != nil {
		log.Print("unmarshaling error: ", err)
		return []interface{}{proto.SignupFormResponse{Ok: true}}
	}

	rows, _ := db.D.Handle.Query(
		"select count(*) from accounts where email = $1",
		form.Account.Email)

	if rows.Next() {
		var count int
		rows.Scan(&count)
		if count > 0 {
			log.Printf("email check %d", count)
			return []interface{}{proto.SignupFormResponse{Ok: false,
				Reason:  proto.SignupFormResponse_EMAIL_ALERT,
				Message: "email already in use"}}
		}
	}

	if len(strings.TrimSpace(form.Account.Username)) > 0 {
		rows, _ = db.D.Handle.Query(
			"select count(*) from accounts where username = $1",
			form.Account.Username)
		if rows.Next() {
			var count int
			rows.Scan(&count)
			if count > 0 {
				log.Printf("username check %d", count)
				return []interface{}{proto.SignupFormResponse{Ok: false,
					Reason:  proto.SignupFormResponse_USERNAME_ALERT,
					Message: "email already in use"}}
			}
		}
	}

	stmt, err := db.D.Handle.Prepare("insert into accounts (id, fullname, email) values ($1, $2, $3)")
	if err != nil {
		log.Printf("prepare %+v", err)
		return []interface{}{proto.SignupFormResponse{Ok: false}}
	}
	sql_result, err := stmt.Exec(
		db.NewId("accounts"),
		form.Account.Fullname,
		form.Account.Email)
	if err != nil {
		log.Printf("upsert %+v", err)
		return []interface{}{proto.SignupFormResponse{Ok: false}}
	}
	new_id, err := sql_result.LastInsertId()
	log.Printf("new id %s", new_id)

	resp := []interface{}{proto.SignupFormResponse{Ok: true, Token: "abc123"}}
	return resp
}
