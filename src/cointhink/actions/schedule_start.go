package actions

import (
	"log"

	"cointhink/container"
	"cointhink/model/account"
	"cointhink/model/schedule"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
)

func DoScheduleStart(scheduleStart *proto.ScheduleStart, accountId string) []gproto.Message {
	var responses []gproto.Message

	log.Printf("ScheduleStart lookup %s %s", scheduleStart.ScheduleId, accountId)
	_schedule, err := schedule.FindWithAccount(scheduleStart.ScheduleId, accountId)
	if err != nil {
		responses = append(responses, &proto.ScheduleStartResponse{Ok: false, Message: "unknown schedule id"})
	} else {
		log.Printf("schedule found %v", _schedule)
		if _schedule.AccountId != accountId {
			responses = append(responses, &proto.ScheduleStopResponse{Ok: false, Message: "not owner"})
		} else {
			_account, err := account.Find(accountId)
			if err != nil {
				responses = append(responses, &proto.ScheduleStartResponse{Ok: false, Message: err.Error()})
			} else {
				enabled := schedule.EnabledNow(&_schedule)
				if enabled == false {
					log.Printf("ScheduleStart enable out of date.")
					err = schedule.EnableUntilNextMonth(&_schedule, &_account)
					if err != nil {
						responses = append(responses, &proto.ScheduleStartResponse{Ok: false, Message: err.Error()})
					} else {
						enabled = true
					}
				}

				if enabled {
					schedule.UpdateStatus(&_schedule, proto.Schedule_enabled)
					_algorun, err := container.Start(_account, _schedule)
					if err != nil {
						responses = append(responses, &proto.ScheduleStartResponse{Ok: false, Message: err.Error()})
					} else {
						if _schedule.Executor == proto.Schedule_lambda {
							sr := proto.ScheduleRun{Schedule: &_schedule, Run: _algorun}
							responses = append(responses, &proto.ScheduleListPartial{ScheduleRun: &sr})
						}
						responses = append(responses, &proto.ScheduleStartResponse{Ok: true})
					}
				}
			}
		}
	}

	return responses
}
