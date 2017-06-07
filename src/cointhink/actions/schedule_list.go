package actions

import (
	"log"

	"cointhink/model/schedule"
	"cointhink/proto"

	"github.com/golang/protobuf/jsonpb"
	gproto "github.com/golang/protobuf/proto"
)

func DoScheduleList(scheduleList proto.ScheduleList, json string, accountId string) []gproto.Message {
	err := jsonpb.UnmarshalString(json, &scheduleList)
	if err != nil {
		log.Print("unmarshaling error: ", err)
		return []gproto.Message{&proto.ScheduleListResponse{Ok: false}}
	}

	var responses []gproto.Message

	schedules, err := schedule.List(accountId)
	if err != nil {
		responses = append(responses, &proto.ScheduleListResponse{Ok: false, Message: err.Error()})
		return responses
	}

	responses = append(responses, &proto.ScheduleListResponse{Ok: true, Schedules: schedules})
	return responses
}
