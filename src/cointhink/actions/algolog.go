package actions

import "log"

import "cointhink/proto"
import "cointhink/model/algolog"

import gproto "github.com/golang/protobuf/proto"

func DoAlgolog(_algolog *proto.Algolog, accountId string) []gproto.Message {
	var responses []gproto.Message

	algolog.Insert(_algolog)
	log.Printf("log inserted %s", _algolog.Id)
	return responses
}
