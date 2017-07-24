package actions

import (
	"log"

	"cointhink/model"
	"cointhink/model/algorithm"
	"cointhink/model/schedule"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
)

func DoScheduleCreate(scheduleCreate *proto.ScheduleCreate, accountId string) []gproto.Message {
	var responses []gproto.Message

	_, err := model.AccountFind(accountId)
	if err != nil {
	}

	log.Printf("creating schedule for algorithm %s", scheduleCreate.Schedule.AlgorithmId)
	if algorithm.Owns(scheduleCreate.Schedule.AlgorithmId, accountId) {
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
	} else {
		responses = append(responses, &proto.ScheduleStopResponse{Ok: false, Message: "denied: not owner of algorithm"})
	}

	responses = append(responses, &proto.ScheduleCreateResponse{Ok: true})
	return responses
}
