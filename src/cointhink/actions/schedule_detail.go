package actions

import "log"
import "cointhink/proto"
import gproto "github.com/golang/protobuf/proto"
import "cointhink/model/schedule"

func DoScheduleDetail(_scheduleDetail *proto.ScheduleDetail, accountId string) []gproto.Message {
	var responses []gproto.Message

	item, err := schedule.Find(_scheduleDetail.ScheduleId)
	if err != nil {
		log.Printf("algo err %+v", err)
		responses = append(responses, &proto.ScheduleDetailResponse{Ok: false, Message: "Not Found"})
	} else {
		if item.AccountId == accountId {
			responses = append(responses, &proto.ScheduleDetailResponse{Ok: true,
				Schedule: &item})
		} else {
			responses = append(responses, &proto.ScheduleDetailResponse{Ok: false, Message: "No permission"})
		}
	}
	return responses
}
