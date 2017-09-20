package lxd

import "cointhink/proto"
import "cointhink/q"
import "cointhink/httpclients"
import "cointhink/model/token"
import "cointhink/model/schedule"
import "cointhink/model/algorun"

import "log"

var op_q []*q.AccountOperation

func AddOp(msg *q.AccountOperation) {
	log.Printf("lxd add op Type:%v Status:%v Operation:%v",
		msg.Operation.Type, msg.Operation.Status, msg.Operation.Operation)
	op_q = append(op_q, msg)
	WatchOp(msg)
}

func WatchOp(msg *q.AccountOperation) {
	op, err := lxdCallOperation("GET", msg.Operation.Operation+"/wait")
	if err != nil {
		log.Printf("lxd.WatchOp err: %v", err)
		return
	}
	log.Printf("lxd.WatchOp finished %s %s", msg.Algorun.Id, op.Status)

	if op.Status == "error" {
		log.Printf("WatchOp got error, skipping status check")
	} else {
		algoRun, _ := algorun.Find(msg.Algorun.Id)
		schedule, _ := schedule.Find(msg.Algorun.ScheduleId)
		token, _ := token.FindByAccountId(msg.Algorun.AccountId, algoRun.Id)
		lxdStatus, err := Status(msg.Algorun.Id)
		log.Printf("lxd.WatchOp lxd status: id:%s status:%v err:%v", msg.Algorun.Id,
			lxdStatus.Metadata.Status, err)

		var algorun_state proto.Algorun_States
		if lxdStatus.ErrorCode == 404 {
			if algoRun.Status == proto.Algorun_States_name[int32(proto.Algorun_destroying)] {
				algorun_state = proto.Algorun_deleted
				algorun.UpdateStatus(algoRun, algorun_state)
			} else {
				log.Printf("lxd.WatchOp error, lxd 404 on algorun %s in state %s",
					algoRun.Id, algoRun.Status)
			}
		} else {
			if algoRun.Status == proto.Algorun_States_name[int32(proto.Algorun_building)] &&
				lxdStatus.Metadata.Status == "Stopped" {
				algorun_state = proto.Algorun_starting
				algorun.UpdateStatus(algoRun, algorun_state)

				FilePut(algoRun.Id, "/cointhink/script.py", algoRun.Code)
				FilePut(algoRun.Id, "/cointhink/auth.json", "{\"Token\":\""+token.Token+"\", \"AlgorunId\":\""+algoRun.Id+"\"}\n")
				FilePut(algoRun.Id, "/cointhink/settings.json", schedule.InitialState)
				op := Start(algoRun.Id)
				q.LXDOPq <- q.AccountOperation{Algorun: algoRun, Operation: op}
			}
			if algoRun.Status == proto.Algorun_States_name[int32(proto.Algorun_destroying)] &&
				lxdStatus.Metadata.Status == "Stopped" {
				op := Delete(algoRun.Id)
				q.LXDOPq <- q.AccountOperation{Algorun: algoRun, Operation: op}
			}

			if algoRun.Status == proto.Algorun_States_name[int32(proto.Algorun_starting)] &&
				lxdStatus.Metadata.Status == "Running" {
				algorun_state = proto.Algorun_running
				algorun.UpdateStatus(algoRun, algorun_state)
			}
		}

		// alert client
		sr := proto.ScheduleRun{Schedule: &schedule, Run: algoRun}
		if algoRun.Status == proto.Algorun_States_name[int32(proto.Algorun_deleted)] {
			sr.Run = nil
		}
		g := proto.ScheduleListPartial{ScheduleRun: &sr}

		sockets := httpclients.AccountIdToSockets(msg.Algorun.AccountId)
		for _, socket := range sockets {
			q.OUTq <- q.RpcOut{Socket: socket,
				Response: &q.RpcResponse{Msg: &g, Id: q.RpcId()}}
		}

	}
}
