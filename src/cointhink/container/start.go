package container

import "log"
import "errors"

import "cointhink/lxd"
import "cointhink/model/algorun"
import "cointhink/model/algolog"
import "cointhink/model/algorithm"
import "cointhink/model/token"
import "cointhink/proto"
import "cointhink/q"
import "cointhink/config"

func Start(account proto.Account, schedule proto.Schedule) error {
	runs, err := algorun.FindReady(account.Id, schedule.Id)
	if err != nil {
		log.Printf("Start err: %v", err)
		return err
	} else {
		if len(runs) == 0 {
			_algorithm, err := algorithm.Find(schedule.AlgorithmId)
			if err != nil {
				log.Printf("Start err: %v", err)
				return err
			} else {
				image := "container " + config.C.QueryString("lxd.container")
				if schedule.Executor == proto.Schedule_lambda {
					image = "lambda"
				}
				log.Printf("Start: algo ready. launching from %s", image)
				_algorun := proto.Algorun{AccountId: account.Id,
					AlgorithmId: schedule.AlgorithmId,
					ScheduleId:  schedule.Id,
					Status:      proto.Algorun_States_name[int32(proto.Algorun_building)],
					Code:        _algorithm.Code,
					Image:       image}
				if schedule.Executor == proto.Schedule_lambda {
					_algorun.Status = proto.Algorun_States_name[int32(proto.Algorun_running)]
				}
				algorun.Insert(&_algorun)
				token_ := proto.Token{AccountId: account.Id, AlgorunId: _algorun.Id}
				token.Insert(&token_)
				_algolog := proto.Algolog{AlgorunId: _algorun.Id, Event: "launch", Level: "info",
					Message: "launching " + schedule.AlgorithmId + " on " + image}
				algolog.Insert(&_algolog)
				if schedule.Executor == proto.Schedule_container {
					op := lxd.Launch(lxd.Lxc{Name: _algorun.Id,
						Profiles: []string{"cointhink"},
						Source:   lxd.LxcSource{Type: "image", Alias: image}})
					q.LXDOPq <- q.AccountOperation{Algorun: &_algorun, Operation: op}
				}
			}
		} else {
			log.Printf("Start aborted. existing algoruns %v", runs)
			return errors.New("existing algorun")
		}
	}
	return nil
}
