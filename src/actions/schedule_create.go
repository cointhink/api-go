package actions

import (
	"log"

	"db"
	"proto"

	"github.com/golang/protobuf/jsonpb"
)

func DoScheduleCreate(scheduleCreate proto.ScheduleCreate, json string) []interface{} {
	err := jsonpb.UnmarshalString(json, &scheduleCreate)
	if err != nil {
		log.Print("unmarshaling error: ", err)
		return []interface{}{proto.SessionCreateResponse{Ok: false}}
	}

	var responses []interface{}

	_, err = db.D.Handle.Query("select account_id from tokens where token = $1",
		scheduleCreate.Schedule.AlgorithmId)
	if err != nil {
		log.Print("token sql error: ", err)
		responses = append(responses, proto.SessionCreateResponse{Ok: false})
	}

	return responses
}
