package actions

import (
	"log"

	"mailer"
	"model"
	"proto"
	"token"

	"github.com/golang/protobuf/jsonpb"
)

func DoSigninEmail(msg proto.SigninEmail, json string) []interface{} {
	resp := []interface{}{}

	err := jsonpb.UnmarshalString(json, &msg)
	if err != nil {
		log.Print("unmarshaling error: ", err)
		return resp
	}

	log.Printf("account lookup %s", msg.Email)
	account_id, err := model.AccountFindByEmail(msg.Email)
	if err != nil {
		return resp
	}

	token, err := token.Find(account_id)

	mailer.MailToken(msg.Email, token)
	return resp
}
