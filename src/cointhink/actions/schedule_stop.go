package actions

import (
	"log"

	"cointhink/container"
	"cointhink/model/algorun"
	"cointhink/model/schedule"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
)

func DoScheduleStop(scheduleStop *proto.ScheduleStop, accountId string) []gproto.Message {
	var responses []gproto.Message

	log.Printf("ScheduleStop lookup %s %s", scheduleStop.ScheduleId, accountId)
	_schedule, err := schedule.FindWithAccount(scheduleStop.ScheduleId, accountId)
	if err != nil {
		responses = append(responses, &proto.ScheduleStopResponse{Ok: false, Message: "unknown schedule id"})
	} else {
		log.Printf("%v", _schedule)
		if _schedule.AccountId != accountId {
			return []gproto.Message{&proto.ScheduleStopResponse{Ok: false, Message: "not owner"}}
		}

		schedule.UpdateStatus(&_schedule, proto.Schedule_disabled)

		// move
		_algoruns, err := algorun.FindReady(accountId, _schedule.Id)
		if err != nil {
			log.Print("scheduleStop findReady: ", err)
			responses = append(responses, &proto.ScheduleStartResponse{Ok: false, Message: "unknown status"})
		} else {
			if len(_algoruns) > 0 {
				_algorun := _algoruns[0]
				container.Stop(_algorun)
				sr := proto.ScheduleRun{Schedule: &_schedule, Run: _algorun}
				responses = append(responses, &proto.ScheduleListPartial{ScheduleRun: &sr})
			}
		}
	}

	return responses
}
