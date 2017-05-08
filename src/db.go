package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Db struct {
	handle *sql.DB
}

func (db *Db) open(db_url string) error {
	var err error
	db.handle, err = sql.Open("postgres", db_url)
	return err
}

func (db *Db) upsert(thing interface{}) error {
	var err error
	_, err = db.handle.Query("select * from accounts")
	return err
}
