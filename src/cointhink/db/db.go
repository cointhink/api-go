package db

import "github.com/jmoiron/sqlx"
import _ "github.com/lib/pq"
import "github.com/satori/go.uuid"
import "regexp"
import "strings"

type Db struct {
	Handle *sqlx.DB
}

var D Db

func init() {
	D = Db{}
}

func (db *Db) Open(db_url string) error {
	var err error
	db.Handle, err = sqlx.Open("postgres", db_url)
	db.Handle.MapperFunc(camelCase) // leave fields as-is
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

// alternate mapperfunc
func camelCase(name string) string {
	words, _ := regexp.Compile("([A-Z][a-z]+)")
	camelFull := strings.ToLower(words.ReplaceAllString(name, "_$1"))
	camel := strings.TrimPrefix(camelFull, "_")
	return camel
}
