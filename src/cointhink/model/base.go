package model

import "reflect"
import "strings"
import "fmt"

import "cointhink/db"

func SqlFields(thing interface{}) (string, string, string) {
	iFields := []string{}
	s := reflect.TypeOf(thing)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		iFields = append(iFields, db.CamelCase(f.Name))
	}
	sqlFields := []string{}
	for _, colName := range iFields {
		sqlFields = append(sqlFields, ":"+colName)
	}
	table := db.Tabelize(s.Name())
	columns := strings.Join(iFields, ", ")
	fields := strings.Join(sqlFields, ", ")

	fmt.Printf("%v %v %v\n", table, columns, fields)
	//rows, err := db.D.Handle.Query("describe " + table)
	return table, columns, fields
}

// unfinished ORM
func find(thing interface{}, id string) error {
	table, columns, _ := SqlFields(thing)
	sql := "select " + columns + ", created_at from " + table + " where id = $1"
	// type.New() object here
	err := db.D.Handle.Get(&thing, sql, id)
	if err != nil {
		return err
	} else {
		return nil
	}
}
