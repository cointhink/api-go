package actions

import (
	"log"

	"cointhink/model"
	"cointhink/model/schedule"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
)

func DoScheduleCreate(scheduleCreate *proto.ScheduleCreate, accountId string) []gproto.Message {
	var responses []gproto.Message

	_, err := model.AccountFind(accountId)
	if err != nil {
	}

	_schedule := proto.Schedule{AccountId: accountId,
		AlgorithmId:  scheduleCreate.Schedule.AlgorithmId,
		Status:       proto.Schedule_disabled,
		InitialState: scheduleCreate.Schedule.InitialState}
	log.Printf("inserting sched state %v", _schedule.Status)
	err = schedule.Insert(&_schedule)
	if err != nil {
		responses = append(responses, &proto.ScheduleCreateResponse{Ok: false, Message: err.Error()})
		return responses
	}

	responses = append(responses, &proto.ScheduleCreateResponse{Ok: true})
	return responses
}
