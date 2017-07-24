package actions

import (
	"log"
	"strings"

	"cointhink/db"
	"cointhink/mailer"
	"cointhink/model/account"
	"cointhink/proto"
	"cointhink/token"
	"cointhink/validate"

	gproto "github.com/golang/protobuf/proto"
)

func DoSignupform(form *proto.SignupForm) []gproto.Message {
	rows, _ := db.D.Handle.Query(
		"select count(*) from accounts where email = $1",
		form.Account.Email)

	emailGood, emailFailReason := validate.Email(form.Account.Email)
	if emailGood == false {
		return []gproto.Message{&proto.SignupFormResponse{Ok: false,
			Reason:  proto.SignupFormResponse_EMAIL_ALERT,
			Message: emailFailReason}}
	}

	if rows.Next() {
		var count int
		rows.Scan(&count)
		if count > 0 {
			log.Printf("email check %d", count)
			return []gproto.Message{&proto.SignupFormResponse{Ok: false,
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
				return []gproto.Message{&proto.SignupFormResponse{Ok: false,
					Reason:  proto.SignupFormResponse_USERNAME_ALERT,
					Message: "email already in use"}}
			}
		}
	}

	err := account.Insert(form.Account)
	if err != nil {
		log.Printf("insert %+v", err)
		return []gproto.Message{&proto.SignupFormResponse{Ok: false}}
	} else {
		token := token.InsertToken(form.Account.Id)
		mailer.MailToken(token, form.Account.Email)
		return []gproto.Message{&proto.SignupFormResponse{Ok: true, Token: token}}
	}
}
