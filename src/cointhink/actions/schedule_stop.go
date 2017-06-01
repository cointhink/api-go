package actions

import (
	"log"

	"cointhink/proto"

	"github.com/golang/protobuf/jsonpb"
)

func DoScheduleStop(scheduleStop proto.ScheduleStop, json string, accountId string) []interface{} {
	err := jsonpb.UnmarshalString(json, &scheduleStop)
	if err != nil {
		log.Print("unmarshaling error: ", err)
		return []interface{}{} //nothing
	}

	var responses []interface{}

	return responses
}
