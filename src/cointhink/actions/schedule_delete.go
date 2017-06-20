package actions

import "cointhink/model/schedule"
import "cointhink/model/algorun"
import "cointhink/container"
import "cointhink/proto"

import gproto "github.com/golang/protobuf/proto"

func DoScheduleDelete(scheduleDelete *proto.ScheduleDelete, accountId string) []gproto.Message {
	var responses []gproto.Message

	_schedule, err := schedule.FindWithAccount(scheduleDelete.ScheduleId, accountId)
	if err != nil {
		responses = append(responses, &proto.ScheduleDeleteResponse{Ok: false, Message: "Unknown schedule"})
	} else {
		// best-effort algorun stop
		_algorun, err := algorun.FindFromSchedule(accountId, _schedule.Id)
		if err != nil {
		} else {
			container.Stop(_algorun)
		}
		schedule.UpdateStatus(_schedule, proto.Schedule_deleted)
		_schedule.Status = proto.Schedule_deleted //cheating
		responses = append(responses, &proto.ScheduleDeleteResponse{Ok: true})
		responses = append(responses, &proto.ScheduleListPartial{ScheduleRun: &proto.ScheduleRun{Schedule: &_schedule}})
	}

	return responses
}
