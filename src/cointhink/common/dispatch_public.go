package common

import (
	"log"

	"cointhink/actions"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

func DispatchPublic(class string, object *any.Any) []gproto.Message {
	log.Printf("*- dispatch-public %#v %#v", class, object)
	var resp []gproto.Message
	switch class {
	case "SignupForm":
		it := proto.SignupForm{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoSignupform(&it)
	case "SessionCreate":
		it := proto.SessionCreate{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoSessionCreate(&it)
	case "SigninEmail":
		it := proto.SigninEmail{}
		ptypes.UnmarshalAny(object, &it)
		resp = actions.DoSigninEmail(&it)
	default:
		log.Printf("unknown public method: %s", class)
	}
	return resp
}
