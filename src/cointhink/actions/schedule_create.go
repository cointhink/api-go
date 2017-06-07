package actions

import (
	"log"

	"cointhink/model"
	"cointhink/model/schedule"
	"cointhink/proto"

	"github.com/golang/protobuf/jsonpb"
	gproto "github.com/golang/protobuf/proto"
)

func DoScheduleCreate(scheduleCreate proto.ScheduleCreate, json string, accountId string) []gproto.Message {
	err := jsonpb.UnmarshalString(json, &scheduleCreate)
	if err != nil {
		log.Print("unmarshaling error: ", err)
		return []gproto.Message{&proto.ScheduleCreateResponse{Ok: false}}
	}

	var responses []gproto.Message

	_, err = model.AccountFind(accountId)
	if err != nil {
	}

	_schedule := proto.Schedule{AccountId: accountId,
		AlgorithmId:  scheduleCreate.Schedule.AlgorithmId,
		Status:       proto.Schedule_stopped,
		InitialState: scheduleCreate.Schedule.InitialState}
	err = schedule.Insert(&_schedule)
	if err != nil {
		responses = append(responses, &proto.ScheduleCreateResponse{Ok: false, Message: err.Error()})
		return responses
	}

	responses = append(responses, &proto.ScheduleCreateResponse{Ok: true})
	return responses
}
