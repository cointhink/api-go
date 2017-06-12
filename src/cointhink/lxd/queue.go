package lxd

import "log"

var LXDOPq chan OperationResponse

func WatchOp(msg OperationResponse) {
	log.Printf("WatchOp Type %v Status %v Operation %v", msg.Type, msg.Status, msg.Operation)
}
