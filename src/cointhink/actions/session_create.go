package actions

import (
	"log"

	"cointhink/db"
	"cointhink/model"
	"cointhink/proto"

	"github.com/golang/protobuf/jsonpb"
)

func DoSessionCreate(sessionCreate proto.SessionCreate, json string) []interface{} {
	err := jsonpb.UnmarshalString(json, &sessionCreate)
	if err != nil {
		log.Print("unmarshaling error: ", err)
		return []interface{}{proto.SessionCreateResponse{Ok: false}}
	}

	var responses []interface{}

	accountId, err := model.TokenFindAccountId(sessionCreate.Token)
	if err != nil {
		log.Printf("Bad token %#v %v", sessionCreate.Token, err)
		responses = append(responses, proto.SessionCreateResponse{Ok: false})
	}

	rows, err := db.D.Handle.Query("select fullname, email from accounts where id = $1", accountId)
	if err != nil {
		log.Print("token sql error: ", err)
		responses = append(responses, proto.SessionCreateResponse{Ok: false})
	} else {
		if rows.Next() {
			var fullname string
			var email string
			rows.Scan(&fullname, &email)
			log.Printf("Token good for Account %#v %#v", fullname, email)
			acct := proto.Account{Fullname: fullname, Email: email}
			responses = append(responses, proto.SessionCreateResponse{Ok: true, Account: &acct})
		} else {
			log.Printf("Token has no Account %#v", rows)
			responses = append(responses, proto.SessionCreateResponse{Ok: false})
		}
	}
	rows.Close()
	return responses
}
