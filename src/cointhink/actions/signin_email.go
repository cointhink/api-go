package actions

import (
	"log"

	"cointhink/mailer"
	"cointhink/model/account"
	"cointhink/model/token"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
)

func DoSigninEmail(msg *proto.SigninEmail) []gproto.Message {
	responses := []gproto.Message{}

	log.Printf("account lookup %s", msg.Email)
	account, err := account.FindByEmail(msg.Email)
	if err != nil {
		log.Printf("account lookup err %#v", err)
		responses = append(responses, &proto.SigninEmailResponse{Ok: false, Message: "email not found"})
	} else {
		token_, err := token.FindByAccountId(account.Id, "")
		if err != nil {
			log.Printf("account has no token.")
		} else {
			mailer.MailToken(msg.Email, token_.Token)
			responses = append(responses, &proto.SigninEmailResponse{Ok: true, Message: "email sent."})
		}
	}
	return responses
}
