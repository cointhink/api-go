package cointhink

import (
	"encoding/json"
	"log"

	"actions"
	"proto"
)

func Dispatch(class string, object interface{}) []interface{} {
	log.Printf("dispatch %s %+v", class, object)
	object_json, _ := json.Marshal(object)
	json := string(object_json)
	var resp []interface{}
	switch class {
	case "SignupForm":
		resp = actions.DoSignupform(proto.SignupForm{}, json)
	}
	return resp
}
