package actions

import (
	"log"

	"cointhink/mailer"
	"cointhink/model"
	"cointhink/proto"
	"cointhink/token"

	gproto "github.com/golang/protobuf/proto"
)

func DoSigninEmail(msg *proto.SigninEmail) []gproto.Message {
	resp := []gproto.Message{}

	log.Printf("account lookup %s", msg.Email)
	account_id, err := model.AccountFindByEmail(msg.Email)
	if err != nil {
		log.Printf("account lookup err %#v", err)
		errResp := []gproto.Message{&proto.SigninEmailResponse{Ok: false, Message: "email not found"}}
		return errResp
	}

	token, err := token.Find(account_id)

	mailer.MailToken(msg.Email, token)
	return resp
}
