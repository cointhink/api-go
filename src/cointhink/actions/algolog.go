package actions

import "log"

import "cointhink/proto"
import "cointhink/q"
import "cointhink/httpclients"
import "cointhink/model/algolog"
import "cointhink/model/algorun"

import gproto "github.com/golang/protobuf/proto"

func DoAlgolog(_algolog *proto.Algolog, accountId string) []gproto.Message {
	var responses []gproto.Message

	if algorun.Owns(_algolog.AlgorunId, accountId) {
		algolog.Insert(_algolog)
		log.Printf("log inserted %s", _algolog.Id)
		// load back the same row, for create date
		log, err := algolog.Find(_algolog.Id)
		if(err != nil) {
			socket := httpclients.AccountIdToSocket(accountId)
			if socket != nil {
				q.OUTq <- q.RpcOut{Socket: socket,
					Response: &q.RpcResponse{Msg: log, Id: q.RpcId()}}
			}
		}
	} else {
		log.Printf("algolog ownership failed for %s %s", _algolog.AlgorunId, accountId)
	}
	return responses
}
