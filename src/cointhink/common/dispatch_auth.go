package common

import (
	"encoding/json"
	"log"

	"cointhink/actions"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
)

func DispatchAuth(class string, object interface{}, accountId string) []gproto.Message {
	log.Printf("*- dispatch-auth %#v %#v %#v", class, object, accountId)
	object_json, _ := json.Marshal(object)
	json := string(object_json)
	var resp []gproto.Message
	switch class {
	case "ScheduleCreate":
		resp = actions.DoScheduleCreate(proto.ScheduleCreate{}, json, accountId)
	case "ScheduleList":
		resp = actions.DoScheduleList(proto.ScheduleList{}, json, accountId)
	case "ScheduleStart":
		resp = actions.DoScheduleStart(proto.ScheduleStart{}, json, accountId)
	case "ScheduleStop":
		resp = actions.DoScheduleStop(proto.ScheduleStop{}, json, accountId)
	default:
		log.Printf("unknown private method: %s", class)
	}
	return resp
}
