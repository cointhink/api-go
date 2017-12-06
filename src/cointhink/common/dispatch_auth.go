package common

import (
	"cointhink/actions"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

func DispatchAuth(class string, object *any.Any, token *proto.Token) []gproto.Message {
	var resp []gproto.Message
	switch class {
	case "ScheduleCreate":
		it := proto.ScheduleCreate{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoScheduleCreate(&it, token.AccountId)
	case "ScheduleList":
		it := proto.ScheduleList{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoScheduleList(&it, token.AccountId)
	case "ScheduleStart":
		it := proto.ScheduleStart{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoScheduleStart(&it, token.AccountId)
	case "ScheduleStop":
		it := proto.ScheduleStop{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoScheduleStop(&it, token.AccountId)
	case "ScheduleDelete":
		it := proto.ScheduleDelete{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoScheduleDelete(&it, token.AccountId)
	case "ScheduleDetail":
		it := proto.ScheduleDetail{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoScheduleDetail(&it, token.AccountId)
	case "Algolog":
		it := proto.Algolog{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoAlgolog(&it, token)
	case "TradeSignal":
		it := proto.TradeSignal{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoTradeSignal(&it, token.AccountId)
	case "AlgologList":
		it := proto.AlgologList{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoAlgologList(&it, token.AccountId)
	case "AlgorithmList":
		it := proto.AlgorithmList{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoAlgorithmList(&it, token.AccountId)
	case "AlgorithmDetail":
		it := proto.AlgorithmDetail{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoAlgorithmDetail(&it, token.AccountId)
	case "Notify":
		it := proto.Notify{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoNotify(&it, token.AccountId)
	default:
		//log.Printf("unknown private method: %s", class)
	}
	return resp
}
