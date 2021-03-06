package container

import "log"
import "cointhink/model/algorun"
import "cointhink/model/schedule"
import "cointhink/lxd"
import "cointhink/q"
import "cointhink/proto"

func SyncAll() {
	runs, err := algorun.List()
	if err != nil {
	}

	log.Printf("*Syncing %d algoruns", len(runs))
	for _, run := range runs {
		Sync(run)
	}
}

func Sync(run *proto.Algorun) {
	_schedule, err := schedule.Find(run.ScheduleId)
	if err != nil {
		log.Printf("algorun orphan (no schedule) %s", run.Id)
	} else {
		if _schedule.Executor != proto.Schedule_lambda {
			response, err := lxd.Status(run.Id)
			if err != nil {
				log.Printf("sync err %v %v", run.Id, err)
			} else {
				if !(run.Status == proto.Algorun_States_name[int32(proto.Algorun_deleted)] &&
					response.ErrorCode == 404) {
					log.Printf("container.Sync: RunId:%v Executor:%s RunStatus:%s LxdStatus:%s LxdErr:%d",
						run.Id, _schedule.Executor, run.Status, response.Metadata.Status,
						response.ErrorCode)
				}
				if response.ErrorCode == 404 {
					algorun.UpdateStatus(run, proto.Algorun_deleted)
				} else if response.StatusCode == 200 {
					if response.Metadata.Status == "Stopped" &&
						run.Status == proto.Algorun_States_name[int32(proto.Algorun_destroying)] {
						log.Printf("container.Sync deleting %s", run.Id)
						op := lxd.Delete(run.Id)
						log.Printf("container.Sync pushing onto LXD Q")
						q.LXDOPq <- q.AccountOperation{Algorun: run, Operation: op}
						log.Printf("container.Sync pushed onto LXD Q")
					}
				}
			}
		}
	}
}
