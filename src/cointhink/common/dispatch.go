package cointhink

import (
	"encoding/json"
	"log"

	"cointhink/actions"
	"cointhink/proto"
)

func Dispatch(class string, object interface{}) []interface{} {
	log.Printf("*- dispatch %#v %#v", class, object)
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
	case "ScheduleCreate":
		resp = actions.DoScheduleCreate(proto.ScheduleCreate{}, json)
	default:
		log.Printf("unknown method: %s", class)
	}
	return resp
}
