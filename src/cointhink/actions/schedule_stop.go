package actions

import (
	"log"

	"cointhink/proto"

	"github.com/golang/protobuf/jsonpb"
	gproto "github.com/golang/protobuf/proto"
)

func DoScheduleStop(scheduleStop proto.ScheduleStop, json string, accountId string) []gproto.Message {
	err := jsonpb.UnmarshalString(json, &scheduleStop)
	if err != nil {
		log.Print("unmarshaling error: ", err)
		return []gproto.Message{} //nothing
	}

	var responses []gproto.Message

	return responses
}
