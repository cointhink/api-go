package container

import "log"

import "cointhink/lxd"
import "cointhink/model/algorun"
import "cointhink/proto"

func Start(accountId string, schedule proto.Schedule) error {
	err := algorun.ReadyToLaunch(accountId, schedule.Id)
	if err != nil {
		log.Printf("Start err: %v", err)
		return err
	} else {
		log.Printf("Start: algo ready. launching")
		_algorun := proto.Algorun{AccountId: accountId,
			AlgorithmId: schedule.AlgorithmId,
			ScheduleId:  schedule.Id}
		algorun.Insert(&_algorun)
		lxd.Launch(lxd.Lxc{Name: _algorun.Id,
			Source: lxd.LxcSource{Type: "image", Fingerprint: "6978077ac9f8"}})
	}
	return nil
}
