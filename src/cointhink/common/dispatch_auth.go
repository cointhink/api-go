package common

import (
	"cointhink/actions"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

func DispatchAuth(class string, object *any.Any, accountId string) []gproto.Message {
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
	case "ScheduleDetail":
		it := proto.ScheduleDetail{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoScheduleDetail(&it, accountId)
	case "Algolog":
		it := proto.Algolog{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoAlgolog(&it, accountId)
	case "TradeSignal":
		it := proto.TradeSignal{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoTradeSignal(&it, accountId)
	case "AlgologList":
		it := proto.AlgologList{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoAlgologList(&it, accountId)
	case "AlgorithmList":
		it := proto.AlgorithmList{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoAlgorithmList(&it, accountId)
	case "AlgorithmDetail":
		it := proto.AlgorithmDetail{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoAlgorithmDetail(&it, accountId)
	case "Notify":
		it := proto.Notify{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoNotify(&it, accountId)
	default:
		//log.Printf("unknown private method: %s", class)
	}
	return resp
}
