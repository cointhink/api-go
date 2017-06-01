package common

import (
	"encoding/json"
	"log"

	"cointhink/actions"
	"cointhink/proto"
)

func DispatchPublic(class string, object interface{}) []interface{} {
	log.Printf("*- dispatch-public %#v %#v", class, object)
	object_json, _ := json.Marshal(object)
	json := string(object_json)
	var resp []interface{}
	switch class {
	case "SignupForm":
		resp = actions.DoSignupform(proto.SignupForm{}, json)
	case "SessionCreate":
		resp = actions.DoSessionCreate(proto.SessionCreate{}, json)
	case "SigninEmail":
		resp = actions.DoSigninEmail(proto.SigninEmail{}, json)
	default:
		log.Printf("unknown public method: %s", class)
	}
	return resp
}
