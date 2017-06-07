package actions

import (
	"io/ioutil"
	"log"

	"cointhink/container"
	"cointhink/lxd"
	"cointhink/model/schedule"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
)

func DoScheduleStart(scheduleStart *proto.ScheduleStart, accountId string) []gproto.Message {
	var responses []gproto.Message

	log.Printf("ScheduleStart lookup %s %s", scheduleStart.ScheduleId, accountId)
	schedule, err := schedule.Find(scheduleStart.ScheduleId, accountId)
	if err != nil {
		responses = append(responses, &proto.ScheduleStartResponse{Ok: false, Message: "unknown schedule id"})
	} else {
		log.Printf("schedule found %v", schedule)
		if schedule.AccountId != accountId {
			return []gproto.Message{&proto.ScheduleStopResponse{Ok: false, Message: "not owner"}}
		}

		resp, err := lxd.Status(schedule.Id)
		if err != nil {
			log.Print("LxdStatus: ", err)
			responses = append(responses, &proto.ScheduleStartResponse{Ok: false, Message: "unknown status"})
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
