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
			return []gproto.Message{&proto.ScheduleStopResponse{Ok: false, Message: "not owner"}}
		}

		schedule.UpdateStatus(_schedule, proto.Schedule_running)
		_account, err := account.Find(accountId)
		if err != nil {
			responses = append(responses, &proto.ScheduleStartResponse{Ok: false, Message: err.Error()})
		} else {
			err = container.Start(_account, _schedule)
			if err != nil {
				responses = append(responses, &proto.ScheduleStartResponse{Ok: false, Message: err.Error()})
			}
		}
	}

	return responses
}
