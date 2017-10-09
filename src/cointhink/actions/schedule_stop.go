package actions

import (
	"log"

	"cointhink/container"
	"cointhink/lxd"
	"cointhink/model/algorun"
	"cointhink/model/schedule"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
)

func DoScheduleStop(scheduleStop *proto.ScheduleStop, accountId string) []gproto.Message {
	var responses []gproto.Message

	log.Printf("ScheduleStop lookup %s %s", scheduleStop.ScheduleId, accountId)
	_schedule, err := schedule.FindWithAccount(scheduleStop.ScheduleId, accountId)
	if err != nil {
		responses = append(responses, &proto.ScheduleStopResponse{Ok: false, Message: "unknown schedule id"})
	} else {
		log.Printf("%v", _schedule)
		if _schedule.AccountId != accountId {
			return []gproto.Message{&proto.ScheduleStopResponse{Ok: false, Message: "not owner"}}
		}

		schedule.UpdateStatus(&_schedule, proto.Schedule_disabled)

		// move
		boxes, err := algorun.FindReady(accountId, _schedule.Id)
		if err != nil {
			log.Print("LxdStatus: ", err)
			responses = append(responses, &proto.ScheduleStartResponse{Ok: false, Message: "unknown status"})
		} else {
			if len(boxes) > 0 {
				box := boxes[0]
				stat, err := lxd.Status(box.Id)
				if err != nil {
					log.Print("LxdStatus: ", err)
					responses = append(responses, &proto.ScheduleStartResponse{Ok: false, Message: "unknown status"})
				}
				log.Printf("ScheduleStop LxdStatus errCode:%d status:%v", stat.ErrorCode, stat.Metadata.Status)
				algorun.UpdateStatus(box, proto.Algorun_deleted)
				if stat.ErrorCode != 404 {
					container.Stop(box)
				}
			}
		}
	}

	return responses
}
