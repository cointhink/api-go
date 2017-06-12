package lxd

import "log"

var LXDOPq chan OperationResponse

var op_q []OperationResponse

func AddOp(msg OperationResponse) {
	log.Printf("lxd ADD Type %v Status %v Operation %v", msg.Type, msg.Status, msg.Operation)
	op_q = append(op_q, msg)
	WatchOp(msg)
}

func WatchOp(msg OperationResponse) {
	op, err := lxdCallOperation("GET", msg.Operation+"/wait")
	if err != nil {
		log.Printf("lxd WATCH err: %v", err)
	}
	if op.Status == "Success" {
	}
}
