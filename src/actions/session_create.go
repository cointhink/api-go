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

	var responses []interface{}
	rows, err := db.D.Handle.Query("select account_id from tokens where token = $1", sessionCreate.Token)
	if err != nil {
		log.Print("token sql error: ", err)
		responses = append(responses, proto.SessionCreateResponse{Ok: false})
	}

	if rows.Next() {
		var accountId string
		rows.Scan(&accountId)
		rows2, err := db.D.Handle.Query("select fullname, email from accounts where id = $1", accountId)
		if err != nil {
			log.Print("token sql error: ", err)
			responses = append(responses, proto.SessionCreateResponse{Ok: false})
		} else {
			if rows2.Next() {
				var fullname string
				var email string
				rows2.Scan(&fullname, &email)
				log.Printf("Token good for Account %#v %#v", fullname, email)
				acct := proto.Account{Fullname: fullname, Email: email}
				responses = append(responses, proto.SessionCreateResponse{Ok: true, Account: &acct})
			} else {
				log.Printf("Token has no Account %#v", rows2)
				responses = append(responses, proto.SessionCreateResponse{Ok: false})
			}
		}
	} else {
		log.Printf("Bad token %#v", sessionCreate.Token)
		responses = append(responses, proto.SessionCreateResponse{Ok: false})
	}
	rows.Close()
	return responses
}
