package actions

import "log"

import "cointhink/mailer"
import "cointhink/proto"
import "cointhink/model/account"

import gproto "github.com/golang/protobuf/proto"

func DoNotify(notify *proto.Notify, accountId string) []gproto.Message {
	resp := []gproto.Message{}

	account_, err := account.Find(accountId)
	if err != nil {
		resp = append(resp, &proto.NotifyResponse{Ok: false, ErrMessage: "bad accountId"})
	} else {
		log.Printf("%+v", notify)
		notify.Recipient = account_.Email
		mailer.MailNotify(notify)
		resp = append(resp, &proto.NotifyResponse{Ok: true})
	}
	return resp
}
