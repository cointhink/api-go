package actions

import (
	"cointhink/model/algorun"
	"cointhink/model/schedule"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
	"log"
)

func DoScheduleList(scheduleList *proto.ScheduleList, accountId string) []gproto.Message {
	var responses []gproto.Message

	schedules, err := schedule.List(accountId)

	scheduleRuns := make([]*proto.ScheduleRun, len(schedules))
	for i, schedule := range schedules {
		runs, err := algorun.FindReady(accountId, schedule.Id)
		if err != nil {
			log.Printf("ScheduleList err %v", err)
		}
		var run *proto.Algorun
		if len(runs) > 0 {
			run = runs[0]
		}
		scheduleRuns[i] = &proto.ScheduleRun{Schedule: schedule, Run: run}
	}
	if err != nil {
		responses = append(responses, &proto.ScheduleListResponse{Ok: false, Message: err.Error()})
		return responses
	}

	responses = append(responses, &proto.ScheduleListResponse{Ok: true, Schedules: scheduleRuns})
	return responses
}
