package lxd

import "cointhink/proto"

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

		_ = proto.ScheduleListPartial{}

		//common.OUTq <- RpcOut{socket: msg.socket,
		//	response: &RpcResponse{msg: &g, id: common.RpcId()}}

	}
}
