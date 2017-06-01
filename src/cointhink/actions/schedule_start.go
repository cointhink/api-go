package actions

import (
	"log"

	"cointhink/net"
	"cointhink/proto"

	"github.com/golang/protobuf/jsonpb"
)

func DoScheduleStart(scheduleStart proto.ScheduleStart, json string, accountId string) []interface{} {
	err := jsonpb.UnmarshalString(json, &scheduleStart)
	if err != nil {
		log.Print("unmarshaling error: ", err)
		return []interface{}{} //nothing
	}

	var responses []interface{}

	resp, err := net.LxdStatus("c1")
	if err != nil {
		log.Print("LxdStatus: ", err)
		return []interface{}{} //nothing
	}
	log.Printf("%v", resp)

	return responses
}
