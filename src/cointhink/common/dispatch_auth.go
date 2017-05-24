package common

import (
	"encoding/json"
	"log"

	"cointhink/actions"
	"cointhink/proto"
)

func DispatchAuth(class string, object interface{}, accountId string) []interface{} {
	log.Printf("*- dispatch-auth %#v %#v %#v", class, object, accountId)
	object_json, _ := json.Marshal(object)
	json := string(object_json)
	var resp []interface{}
	switch class {
	case "ScheduleCreate":
		resp = actions.DoScheduleCreate(proto.ScheduleCreate{}, json, accountId)
	case "ScheduleList":
		resp = actions.DoScheduleList(proto.ScheduleList{}, json, accountId)
	default:
		log.Printf("unknown method: %s", class)
	}
	return resp
}
