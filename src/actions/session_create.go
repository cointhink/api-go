package actions

import (
	"log"

	"db"
	"proto"

	"github.com/golang/protobuf/jsonpb"
)

func DoSessionCreate(sessionCreate proto.SessionCreate, json string) []interface{} {
	err := jsonpb.UnmarshalString(json, &sessionCreate)
	if err != nil {
		log.Print("unmarshaling error: ", err)
		return []interface{}{proto.SessionCreateResponse{Ok: false}}
	}
	rows, err := db.D.Handle.Query("select account_id from tokens where token = $1", sessionCreate.Token)
	if err != nil {
		log.Print("token sql error: ", err)
		return []interface{}{proto.SessionCreateResponse{Ok: false}}
	}

	if rows.Next() {
		log.Printf("toekn lookup got next")

		var accountId string
		rows.Scan(&accountId)
		rows2, err := db.D.Handle.Query("select fullname, email from accounts where id = $1", accountId)
		if err != nil {
			log.Print("token sql error: ", err)
			return []interface{}{proto.SessionCreateResponse{Ok: false}}
		}
		if rows2.Next() {
			var fullname string
			var email string
			rows2.Scan(&fullname, &email)
			log.Printf("account found %s %s", fullname, email)
			acct := proto.Account{Fullname: fullname, Email: email}
			return []interface{}{proto.SessionCreateResponse{Ok: true, Account: &acct}}
		} else {
			return []interface{}{proto.SessionCreateResponse{Ok: true}}
		}
	} else {
		log.Printf("toekn lookup no next", rows)
		return []interface{}{proto.SessionCreateResponse{Ok: false}}
	}
}
