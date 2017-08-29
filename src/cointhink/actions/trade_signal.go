package actions

import "log"

import "cointhink/mailer"
import "cointhink/proto"
import "cointhink/model/account"

import gproto "github.com/golang/protobuf/proto"

func DoTradeSignal(tradeSignal *proto.TradeSignal, accountId string) []gproto.Message {
	resp := []gproto.Message{}

	account_, err := account.Find(accountId)
	if err != nil {
		resp = append(resp, &proto.TradeSignalResponse{Ok: false, Message: "bad accountId"})
	} else {
		log.Printf("%+v", tradeSignal)
		mailer.MailTrade(account_.Email, tradeSignal)
		resp = append(resp, &proto.TradeSignalResponse{Ok: true})
	}
	return resp
}
