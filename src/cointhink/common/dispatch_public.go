package common

import (
	"log"

	"cointhink/actions"
	"cointhink/proto"

	"github.com/golang/protobuf/jsonpb"
	gproto "github.com/golang/protobuf/proto"
)

func DispatchPublic(class string, json string) []gproto.Message {
	log.Printf("*- dispatch-public %#v %#v", class, json)
	var resp []gproto.Message
	switch class {
	case "SignupForm":
		it := proto.SignupForm{}
		jsonpb.UnmarshalString(json, &it)
		resp = actions.DoSignupform(&it)
	case "SessionCreate":
		it := proto.SessionCreate{}
		jsonpb.UnmarshalString(json, &it)
		resp = actions.DoSessionCreate(&it)
	case "SigninEmail":
		it := proto.SigninEmail{}
		jsonpb.UnmarshalString(json, &it)
		resp = actions.DoSigninEmail(&it)
	default:
		log.Printf("unknown public method: %s", class)
	}
	return resp
}
