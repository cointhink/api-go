package common

import (
	"log"

	"cointhink/actions"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

func DispatchAuth(class string, object *any.Any, accountId string) []gproto.Message {
	log.Printf("*- dispatch-auth %#v %s %#v", class, object, accountId)
	var resp []gproto.Message
	switch class {
	case "ScheduleCreate":
		it := proto.ScheduleCreate{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoScheduleCreate(&it, accountId)
	case "ScheduleList":
		it := proto.ScheduleList{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoScheduleList(&it, accountId)
	case "ScheduleStart":
		it := proto.ScheduleStart{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoScheduleStart(&it, accountId)
	case "ScheduleStop":
		it := proto.ScheduleStop{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoScheduleStop(&it, accountId)
	case "ScheduleDelete":
		it := proto.ScheduleDelete{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoScheduleDelete(&it, accountId)
	case "Algolog":
		it := proto.Algolog{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoAlgolog(&it, accountId)
	default:
		log.Printf("unknown private method: %s", class)
	}
	return resp
}
