package actions

import (
	"log"

	"mailer"
	"proto"

	"github.com/golang/protobuf/jsonpb"
)

func DoSigninEmail(form proto.SigninEmail, json string) []interface{} {
	err := jsonpb.UnmarshalString(json, &form)
	if err != nil {
		log.Print("unmarshaling error: ", err)
		return []interface{}{}
	}

	mailer.MailToken("token", "form.Account.Email")
	return []interface{}{}
}
