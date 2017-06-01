package actions

import (
	"log"
	"net/http"
	"time"

	"cointhink/config"
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

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	url := config.C.QueryString("lxd.api_url")
	log.Printf("lxd call %s", url)
	netClient.Get(url)

	return responses
}
