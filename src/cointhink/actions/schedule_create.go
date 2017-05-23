package actions

import (
	"log"

	"cointhink/model"
	"cointhink/proto"

	"github.com/golang/protobuf/jsonpb"
)

func DoScheduleCreate(scheduleCreate proto.ScheduleCreate, json string, accountId string) []interface{} {
	err := jsonpb.UnmarshalString(json, &scheduleCreate)
	if err != nil {
		log.Print("unmarshaling error: ", err)
		return []interface{}{proto.ScheduleCreateResponse{Ok: false}}
	}

	var responses []interface{}

	_, err = model.AccountFind(accountId)
	if err != nil {
	}

	err = model.ScheduleInsert(accountId, scheduleCreate.Schedule.AlgorithmId, "active")
	if err != nil {
		responses = append(responses, proto.ScheduleCreateResponse{Ok: false, Message: err.Error()})
		return responses
	}

	responses = append(responses, proto.ScheduleCreateResponse{Ok: true})
	return responses
}
