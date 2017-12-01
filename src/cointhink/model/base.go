package model

import "reflect"
import "cointhink/db"

func PbFields(thing interface{}) []string {
	var fields = []string{}
	s := reflect.TypeOf(thing)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fields = append(fields, db.CamelCase(f.Name))
	}
	return fields
}
