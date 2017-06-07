package actions

import (
	"log"

	"cointhink/db"
	"cointhink/model"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
)

func DoSessionCreate(sessionCreate *proto.SessionCreate) []gproto.Message {
	var responses []gproto.Message

	accountId, err := model.TokenFindAccountId(sessionCreate.Token)
	if err != nil {
		log.Printf("Bad token %#v %v", sessionCreate.Token, err)
		responses = append(responses, &proto.SessionCreateResponse{Ok: false})
	}

	rows, err := db.D.Handle.Query("select fullname, email from accounts where id = $1", accountId)
	if err != nil {
		log.Print("token sql error: ", err)
		responses = append(responses, &proto.SessionCreateResponse{Ok: false})
	} else {
		if rows.Next() {
			var fullname string
			var email string
			rows.Scan(&fullname, &email)
			log.Printf("Token good for Account %#v %#v", fullname, email)
			acct := proto.Account{Fullname: fullname, Email: email}
			responses = append(responses, &proto.SessionCreateResponse{Ok: true, Account: &acct})
		} else {
			log.Printf("Token has no Account %#v", rows)
			responses = append(responses, &proto.SessionCreateResponse{Ok: false})
		}
	}
	rows.Close()
	return responses
}
