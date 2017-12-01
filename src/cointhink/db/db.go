package db

import "github.com/jmoiron/sqlx"
import _ "github.com/lib/pq"
import "github.com/satori/go.uuid"
import "regexp"
import "strings"
import "fmt"
import "reflect"

type Db struct {
	Handle *sqlx.DB
}

type SqlDetail struct {
	Table   string
	Columns string
	Fields  string
}

var D Db = Db{}
var registeredTables map[string]SqlDetail = map[string]SqlDetail{}

func (db *Db) Open(db_url string) error {
	var err error
	db.Handle, err = sqlx.Open("postgres", db_url)
	db.Handle.MapperFunc(CamelCase) // StructField becomes struct_field in database
	schemaCheck()
	return err
}

type PsqlRow struct {
	TableName  string
	ColumnName string
}

func schemaCheck() {
	for table, _ := range registeredTables {
		psqlRows := []PsqlRow{}
		err := D.Handle.Select(&psqlRows, "SELECT table_name, column_name FROM information_schema.columns WHERE "+
			"table_schema = 'public' AND table_name = $1", table)
		if err != nil {
			fmt.Printf("rows err %v\n", err)
		} else {
			// for _, row := range psqlRows {
			// 	log.Printf("table %s check rows %v\n", table, row.ColumnName)
			// }
		}
	}
}

func NewId(table_name string) string {
	uuid_str := uuid.NewV4().String()
	return table_name + "-" + uuid_str[19:len(uuid_str)]
}

// alternate mapperfunc
func CamelCase(name string) string {
	words, _ := regexp.Compile("([A-Z][a-z]+)")
	camelFull := strings.ToLower(words.ReplaceAllString(name, "_$1"))
	camel := strings.TrimPrefix(camelFull, "_")
	return camel
}

func Tabelize(name string) string {
	return strings.ToLower(name) + "s"
}

func Register(thing interface{}) SqlDetail {
	detail := SqlFields(thing)
	registeredTables[detail.Table] = detail
	fmt.Printf("model registered: %v | %v\n", detail.Table, detail.Columns)
	return detail
}

func SqlFields(thing interface{}) SqlDetail {
	iFields := []string{}
	s := reflect.TypeOf(thing)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		iFields = append(iFields, CamelCase(f.Name))
	}
	sqlFields := []string{}
	for _, colName := range iFields {
		sqlFields = append(sqlFields, ":"+colName)
	}
	table := Tabelize(s.Name())
	columns := strings.Join(iFields, ", ")
	fields := strings.Join(sqlFields, ", ")

	return SqlDetail{Table: table, Columns: columns, Fields: fields}
}
