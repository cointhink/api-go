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
	case "ScheduleCreate":
		resp = actions.DoScheduleCreate(proto.ScheduleCreate{}, json)
	default:
		log.Printf("unknown method: %s", class)
	}
	return resp
}

func DispatchAuth(class string, object interface{}, accountId string) []interface{} {
	log.Printf("*- dispatch-auth %#v %#v %#v", class, object, accountId)
	object_json, _ := json.Marshal(object)
	json := string(object_json)
	var resp []interface{}
	switch class {
	case "ScheduleCreate":
		resp = actions.DoScheduleCreate(proto.ScheduleCreate{}, json)
	default:
		log.Printf("unknown method: %s", class)
	}
	return resp
}
