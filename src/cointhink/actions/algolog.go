package actions

import "log"

import "cointhink/proto"
import "cointhink/model/algolog"
import "cointhink/model/algorun"
import "cointhink/q"
import "cointhink/httpclients"

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
			// FIXME: sent only to web clients
			for _, clientSocket := range httpclients.AccountIdToWebSockets(accountId) {
				q.OUTq <- q.RpcOut{Socket: clientSocket,
					Response: &q.RpcResponse{Msg: log_readback, Id: "algolog-echo"}}
			}
		}
	} else {
		log.Printf("algolog ownership failed for %s %s", _algolog.AlgorunId, accountId)
	}
	return responses
}
