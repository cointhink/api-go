package db

import "database/sql"
import _ "github.com/lib/pq"
import "github.com/satori/go.uuid"

type Db struct {
	Handle *sql.DB
}

var D Db

func init() {
	D = Db{}
}

func (db *Db) Open(db_url string) error {
	var err error
	db.Handle, err = sql.Open("postgres", db_url)
	return err
}

func (db *Db) upsert(thing interface{}) error {
	var err error
	_, err = db.Handle.Query("select * from accounts")
	return err
}

func NewId(table_name string) string {
	uuid_str := uuid.NewV4().String()
	return table_name + "-" + uuid_str[19:len(uuid_str)]
}
