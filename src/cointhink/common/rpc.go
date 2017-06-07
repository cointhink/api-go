package common

import "cointhink/proto"
import gproto "github.com/golang/protobuf/proto"

func rpcClass(name string) gproto.Message {
	switch name {
	case "SessionCreate":
		return &proto.SessionCreate{}
	case "ScheduleList":
		return &proto.ScheduleList{}
	}
	return nil
}
