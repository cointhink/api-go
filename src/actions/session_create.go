package actions

import (
	"log"

	"db"
	"proto"

	"github.com/golang/protobuf/jsonpb"
)

func DoSessionCreate(form proto.SessionCreate, json string) []interface{} {
	err := jsonpb.UnmarshalString(json, &form)
	if err != nil {
		log.Print("unmarshaling error: ", err)
		return []interface{}{proto.SessionCreateResponse{Ok: false}}
	}
	rows, err := db.D.Handle.Query("select * from tokens where token =$1")
	log.Printf("%#v", rows)

	resp := []interface{}{proto.SessionCreateResponse{Ok: true}}
	return resp

}
