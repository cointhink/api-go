package actions

import (
	"io/ioutil"
	"log"

	"cointhink/container"
	"cointhink/model/schedule"
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
	schedule, err := schedule.Find(scheduleStart.ScheduleId, accountId)
	if err != nil {
		responses = append(responses, proto.ScheduleStartResponse{Ok: false, Message: "unknown schedule id"})
	} else {
		log.Printf("schedule found %v", schedule)
		resp, err := net.LxdStatus(schedule.Id)
		if err != nil {
			log.Print("LxdStatus: ", err)
			responses = append(responses, proto.ScheduleStartResponse{Ok: false, Message: "unknown status"})
		}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		log.Printf("LxdStatus %v %v", resp.Status, string(bodyBytes))
		resp.Body.Close()
		if resp.StatusCode == 404 {
			container.Launch(accountId, schedule.AlgorithmId)
		} else {
			log.Printf("container not launched: exists")
		}
	}

	return responses
}
