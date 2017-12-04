package container

import "log"

import "cointhink/lxd"
import "cointhink/q"
import "cointhink/proto"
import "cointhink/model/algorun"
import "cointhink/model/schedule"

func Stop(_algorun *proto.Algorun) error {
	log.Printf("Stop: %s", _algorun.Id)
	_schedule, err := schedule.Find(_algorun.ScheduleId)
	if err != nil {
	} else {
		if _schedule.Executor == proto.Schedule_container ||
			_schedule.Executor == proto.Schedule_lambda_master {
			algorun.UpdateStatus(_algorun, proto.Algorun_destroying)
			op := lxd.Stop(_algorun.Id)
			if op.Type != "error" {
				q.LXDOPq <- q.AccountOperation{Algorun: _algorun, Operation: op}
			}
		} else {
			algorun.UpdateStatus(_algorun, proto.Algorun_deleted)
		}
	}
	return nil
}
