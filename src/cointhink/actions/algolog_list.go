package actions

import "log"
import "cointhink/proto"

import "cointhink/model/algolog"
import "cointhink/model/algorun"

import gproto "github.com/golang/protobuf/proto"

func DoAlgologList(_algologList *proto.AlgologList, accountId string) []gproto.Message {
	var responses []gproto.Message

	if algorun.Owns(_algologList.AlgorunId, accountId) {
		logs, err := algolog.FindAll(_algologList.AlgorunId, 3)
		if err != nil {
			log.Printf("%+v", err)
		} else {
			log.Printf("algorun %s logs count %d", _algologList.AlgorunId, len(logs))
			// append(responses, logs...) Why, go?
			for _, log := range logs {
				responses = append(responses, log)
			}
		}
	} else {
		log.Printf("algorun ownership failed for %s %s", _algologList.AlgorunId, accountId)
	}

	return responses
}
