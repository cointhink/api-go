package actions

import (
	"log"

	"cointhink/model"
	"cointhink/model/account"
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

	_account, err := account.Find(accountId)
	if err != nil {
		log.Print("token sql error: ", err)
		responses = append(responses, &proto.SessionCreateResponse{Ok: false})
	} else {
		log.Printf("Token good for Account %#v %#v", _account.Fullname, _account.Email)
		responses = append(responses, &proto.SessionCreateResponse{Ok: true, Account: &_account})
	}
	return responses
}
