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

func Start(account proto.Account, _schedule proto.Schedule) (*proto.Algorun, error) {
	runs, err := algorun.FindReady(account.Id, _schedule.Id)
	if err != nil {
		log.Printf("Start err: %v", err)
		return nil, err
	} else {
		if len(runs) == 0 {
			_algorithm, err := algorithm.Find(_schedule.AlgorithmId)
			if err != nil {
				log.Printf("Start err: %v", err)
				return nil, err
			} else {
				image := config.C.QueryString("lxd.container")
				if _schedule.Executor == proto.Schedule_lambda {
					image = "lambda"
				}
				log.Printf("Start: algo ready. launching from %s", image)
				_algorun := proto.Algorun{AccountId: account.Id,
					AlgorithmId: _schedule.AlgorithmId,
					ScheduleId:  _schedule.Id,
					Status:      proto.Algorun_States_name[int32(proto.Algorun_building)],
					Code:        _algorithm.Code,
					Image:       image}
				if _schedule.Executor == proto.Schedule_lambda {
					_algorun.Status = proto.Algorun_States_name[int32(proto.Algorun_running)]
				}
				algorun.Insert(&_algorun)
				token_ := proto.Token{AccountId: account.Id, AlgorunId: _algorun.Id}
				token.Insert(&token_)
				_algolog := proto.Algolog{AlgorunId: _algorun.Id, Event: "launch", Level: "info",
					Message: "launching " + _schedule.AlgorithmId + " on " + image}
				algolog.Insert(&_algolog)
				if _schedule.Executor == proto.Schedule_container ||
					_schedule.Executor == proto.Schedule_lambda_master {
					op := lxd.Launch(lxd.Lxc{Name: _algorun.Id,
						Profiles: []string{"cointhink"},
						Source:   lxd.LxcSource{Type: "image", Alias: image}})
					q.LXDOPq <- q.AccountOperation{Algorun: &_algorun, Operation: op}
				}
				return &_algorun, nil
			}
		} else {
			log.Printf("Start aborted. existing algoruns %v", runs)
			return nil, errors.New("existing algorun")
		}
	}
}
