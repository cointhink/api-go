package container

import "log"
import "errors"

import "cointhink/lxd"
import "cointhink/q"
import "cointhink/proto"
import "cointhink/model/algorun"
import "cointhink/model/schedule"

func Stop(_algorun *proto.Algorun) error {
	log.Printf("container stop: %s", _algorun.Id)
	_schedule, err := schedule.Find(_algorun.ScheduleId)
	if err != nil {
		return err
	} else {
		if _schedule.Executor == proto.Schedule_container ||
			_schedule.Executor == proto.Schedule_lambda_master {
			stat, err := lxd.Status(_algorun.Id)
			if err != nil {
				return err
			} else {
				if stat.ErrorCode != 404 {
					algorun.UpdateStatus(_algorun, proto.Algorun_destroying)
					op := lxd.Stop(_algorun.Id)
					if op.Type != "error" {
						q.LXDOPq <- q.AccountOperation{Algorun: _algorun, Operation: op}
						return nil
					} else {
						return errors.New("lxd Stop failed")
					}
				} else {
					algorun.UpdateStatus(_algorun, proto.Algorun_deleted)
					return nil
				}
			}
		} else {
			algorun.UpdateStatus(_algorun, proto.Algorun_deleted)
			return nil
		}
	}
}
