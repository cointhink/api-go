package actions

import "log"

import "cointhink/proto"
import "cointhink/model/algolog"
import "cointhink/model/algorun"

import gproto "github.com/golang/protobuf/proto"

func DoAlgolog(_algolog *proto.Algolog, accountId string) []gproto.Message {
	var responses []gproto.Message

	if algorun.Owns(_algolog.AlgorunId, accountId) {
		algolog.Insert(_algolog)
		log.Printf("log inserted %s", _algolog.Id)
	} else {
		log.Printf("algolog ownership failed for %s %s", _algolog.AlgorunId, accountId)
	}
	return responses
}
