package actions

import (
	"log"

	"cointhink/config"
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

	url := config.C.QueryString("lxd.api_url")
	log.Printf("lxd call %s", url)
	net.Client.Get(url)

	return responses
}
