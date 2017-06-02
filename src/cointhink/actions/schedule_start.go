package actions

import (
	"log"

	"cointhink/model"
	"cointhink/net"
	"cointhink/proto"

	"github.com/golang/protobuf/jsonpb"
)

func DoScheduleStart(scheduleStart proto.ScheduleStart, json string, accountId string) []interface{} {
	err := jsonpb.UnmarshalString(json, &scheduleStart)
	if err != nil {
		log.Print("unmarshaling error: ", err)
		return []interface{}{proto.ScheduleStartResponse{Ok: false}}
	}

	var responses []interface{}

	log.Printf("ScheduleStart lookup %s %s", scheduleStart.ScheduleId, accountId)
	schedule, err := model.ScheduleFind(scheduleStart.ScheduleId, accountId)
	if err != nil {
		responses = append(responses, proto.ScheduleStartResponse{Ok: false, Message: "unknown schedule id"})
	} else {
		log.Printf("schedule found %v", schedule)
		resp, err := net.LxdStatus(schedule.Id)
		if err != nil {
			log.Print("LxdStatus: ", err)
			responses = append(responses, proto.ScheduleStartResponse{Ok: false, Message: "unknown status"})
		}
		log.Printf("%v", resp)
		if resp.StatusCode == 404 {
			net.LxdLaunch(net.Lxc{Name: schedule.Id,
				Source: net.LxcSource{Type: "image", Alias: "alpine"}})
		}
	}

	return responses
}
