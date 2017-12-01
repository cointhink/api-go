package model

import "cointhink/db"

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
