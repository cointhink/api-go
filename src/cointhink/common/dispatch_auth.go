package common

import (
	"log"

	"cointhink/actions"
	"cointhink/proto"

	"github.com/golang/protobuf/jsonpb"
	gproto "github.com/golang/protobuf/proto"
)

func DispatchAuth(class string, json string, accountId string) []gproto.Message {
	log.Printf("*- dispatch-auth %#v %#v %#v", class, json, accountId)
	var resp []gproto.Message
	switch class {
	case "ScheduleCreate":
		it := proto.ScheduleCreate{}
		jsonpb.UnmarshalString(json, &it)
		resp = actions.DoScheduleCreate(&it, accountId)
	case "ScheduleList":
		it := proto.ScheduleList{}
		jsonpb.UnmarshalString(json, &it)
		resp = actions.DoScheduleList(&it, accountId)
	case "ScheduleStart":
		it := proto.ScheduleStart{}
		jsonpb.UnmarshalString(json, &it)
		resp = actions.DoScheduleStart(&it, accountId)
	case "ScheduleStop":
		it := proto.ScheduleStop{}
		jsonpb.UnmarshalString(json, &it)
		resp = actions.DoScheduleStop(&it, accountId)
	default:
		log.Printf("unknown private method: %s", class)
	}
	return resp
}
