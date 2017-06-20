package container

import "log"
import "errors"

import "cointhink/lxd"
import "cointhink/model/algorun"
import "cointhink/proto"

func Start(account proto.Account, schedule proto.Schedule) error {
	runs, err := algorun.FindReady(account.Id, schedule.Id)
	if err != nil {
		log.Printf("Start err: %v", err)
		return err
	} else {
		if len(runs) == 0 {
			log.Printf("Start: algo ready. launching")
			_algorun := proto.Algorun{AccountId: account.Id,
				AlgorithmId: schedule.AlgorithmId,
				ScheduleId:  schedule.Id,
				Status:      proto.Algorun_States_name[int32(proto.Algorun_building)]}
			algorun.Insert(&_algorun)
			op := lxd.Launch(lxd.Lxc{Name: _algorun.Id,
				Source: lxd.LxcSource{Type: "image", Fingerprint: "6978077ac9f8"}})
			lxd.LXDOPq <- lxd.AccountOperation{Algorun: &_algorun, Operation: op}
			//op = lxd.Start(_algorun.Id)
			//op = lxd.Exec(_algorun.Id, "/bin/ls")
		} else {
			log.Printf("Start aborted. existing algoruns %v", runs)
			return errors.New("existing algorun")
		}
	}
	return nil
}
