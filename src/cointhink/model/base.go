package model

import "reflect"
import "strings"

import "cointhink/db"

// unfinished ORM
func SqlFields(thing interface{}) (string, string, string) {
	columns := []string{}
	s := reflect.TypeOf(thing)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		columns = append(columns, db.CamelCase(f.Name))
	}
	sqlFields := []string{}
	for _, colName := range columns {
		sqlFields = append(sqlFields, ":"+colName)
	}
	return db.Tabelize(s.Name()),
		strings.Join(columns, ", "),
		strings.Join(sqlFields, ", ")
}

func find(thing interface{}, id string) error {
	table, columns, _ := SqlFields(thing)
	sql := "select " + columns + ", created_at from " + table + " where id = $1"
	err := db.D.Handle.Get(&thing, sql, id)
	if err != nil {
		return err
	} else {
		return nil
	}
}
