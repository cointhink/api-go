package actions

import (
	"cointhink/model/schedule"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
)

func DoScheduleList(scheduleList *proto.ScheduleList, accountId string) []gproto.Message {
	var responses []gproto.Message

	schedules, err := schedule.List(accountId)
	if err != nil {
		responses = append(responses, &proto.ScheduleListResponse{Ok: false, Message: err.Error()})
		return responses
	}

	responses = append(responses, &proto.ScheduleListResponse{Ok: true, Schedules: schedules})
	return responses
}
