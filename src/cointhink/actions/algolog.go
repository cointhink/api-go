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
		// load back the same row, for create date
		log_readback, err := algolog.Find(_algolog.Id)
		if err != nil {
			log.Printf("algolog readback err %+v", err)
		} else {
			_ = log_readback
			// FIXME: sent only to web clients
			// sockets := httpclients.AccountIdToSockets(accountId)
			// for _, socket := range sockets {
			// 	q.OUTq <- q.RpcOut{Socket: socket,
			// 		Response: &q.RpcResponse{Msg: log_readback, Id: q.RpcId()}}
			// }
		}
	} else {
		log.Printf("algolog ownership failed for %s %s", _algolog.AlgorunId, accountId)
	}
	return responses
}
