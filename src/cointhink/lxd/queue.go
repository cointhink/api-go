package lxd

import "cointhink/proto"
import "cointhink/q"
import "cointhink/httpclients"
import "cointhink/model/schedule"

import "cointhink/model/algorun"

import "log"

import "github.com/google/uuid"

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
	log.Printf("lxd.WatchOp finished %s %s", msg.Algorun.Id, op.Status)
	status, err := Status(msg.Algorun.Id)
	log.Printf("lxd.WatchOp status: id:%s status:%v err:%v", msg.Algorun.Id, status.Metadata.Status, err)

	var algorun_state proto.Algorun_States
	if status.ErrorCode == 404 {
		algorun_state = proto.Algorun_deleted
	} else {

		switch status.Metadata.Status {
		case "Stopped":
			algorun_state = proto.Algorun_stopped
		}
	}

	algorun.UpdateStatus(msg.Algorun, algorun_state)

	s, _ := schedule.Find(msg.Algorun.ScheduleId)
	sr := proto.ScheduleRun{Schedule: &s, Run: msg.Algorun}
	g := proto.ScheduleListPartial{ScheduleRun: &sr}

	socket := httpclients.AccountIdToSocket(msg.Algorun.AccountId)
	log.Printf("Watchop socket lookup id %p", socket)
	q.OUTq <- q.RpcOut{Socket: socket,
		Response: &q.RpcResponse{Msg: &g, Id: RpcId()}}
}

func RpcId() string {
	uuid, _ := uuid.NewRandom()
	uuidStr := uuid.String()
	return uuidStr[19:35]
}
