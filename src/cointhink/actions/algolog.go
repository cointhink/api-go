package actions

import "log"

import "cointhink/proto"
import "cointhink/model/algolog"
import "cointhink/q"
import "cointhink/httpclients"

import gproto "github.com/golang/protobuf/proto"

func DoAlgolog(_algolog *proto.Algolog, token *proto.Token) []gproto.Message {
	var responses []gproto.Message

	_algolog.AlgorunId = token.AlgorunId
	algolog.Insert(_algolog)
	log.Printf("log inserted %s", _algolog.Id)
	// load back the same row, for create date
	log_readback, err := algolog.Find(_algolog.Id)
	if err != nil {
		log.Printf("algolog readback err %+v", err)
	} else {
		for _, clientSocket := range httpclients.AccountIdToWebSockets(token.AccountId) {
			q.OUTq <- q.RpcOut{Socket: clientSocket,
				Response: &q.RpcResponse{Msg: log_readback, Id: "algolog-echo"}}
		}
	}
	return responses
}
