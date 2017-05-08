package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
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

func NewId(table_name string) string {
	uuid_str := uuid.NewV4().String()
	return table_name + "-" + uuid_str[19:len(uuid_str)]
}
