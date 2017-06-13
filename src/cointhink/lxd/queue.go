package lxd

import "cointhink/proto"
import "cointhink/q"
import "cointhink/httpclients"

import "log"

var LXDOPq chan AccountOperation

var op_q []*AccountOperation

func AddOp(msg *AccountOperation) {
	log.Printf("lxd ADD Type %v Status %v Operation %v",
		msg.Operation.Type, msg.Operation.Status, msg.Operation.Operation)
	op_q = append(op_q, msg)
	WatchOp(msg)
}

func WatchOp(msg *AccountOperation) {
	op, err := lxdCallOperation("GET", msg.Operation.Operation+"/wait")
	if err != nil {
		log.Printf("lxd WATCH err: %v", err)
	}
	if op.Status == "Success" {
		log.Printf("op success for %v", msg.Account.Email)

		g := proto.ScheduleListPartial{}

		socket := httpclients.AccountIdToSocket(msg.Account.Id)
		log.Printf("Watchop socket lookup %p", socket)
		q.OUTq <- q.RpcOut{Socket: socket,
			Response: &q.RpcResponse{Msg: &g, Id: "z"}}

	}
}
