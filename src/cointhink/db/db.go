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
	Model            string
	Table            string
	Columns          []string
	ColumnsSql       string
	ColumnsInsertSql string
	FieldsSql        string
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
	for table, registeredSchema := range registeredTables {
		psqlRows := []PsqlRow{}
		err := D.Handle.Select(&psqlRows, "SELECT table_name, column_name FROM information_schema.columns WHERE "+
			"table_schema = 'public' AND table_name = $1", table)
		if err != nil {
			fmt.Printf("rows err %v\n", err)
		} else {
			for _, row := range psqlRows {
				found := false
				for _, registeredColumnName := range registeredSchema.Columns {
					if row.ColumnName == registeredColumnName {
						found = true
					}
				}
				if !found {
					fmt.Printf("WARNING: table %s check psql column %s is not in protobuf\n", table, row.ColumnName)
				}
			}
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
	detail := SqlFieldsSql(thing)
	registeredTables[detail.Table] = detail
	fmt.Printf("model %v | %v - %v\n", detail.Model, detail.Table, detail.ColumnsSql)
	return detail
}

func SqlFieldsSql(thing interface{}) SqlDetail {
	s := reflect.TypeOf(thing)
	model := s.Name()
	table := Tabelize(model)
	iFields := []string{}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		iFields = append(iFields, CamelCase(f.Name))
	}
	sqlFieldsSql := []string{}
	for _, colName := range iFields {
		sqlFieldsSql = append(sqlFieldsSql, ":"+colName)
	}
	tableFields := []string{}
	for _, iField := range iFields {
		tableFields = append(tableFields, table+"."+iField)
	}
	columnsSql := strings.Join(tableFields, ", ")
	columnsInsertSql := strings.Join(iFields, ", ")
	fieldsSql := strings.Join(sqlFieldsSql, ", ")

	return SqlDetail{Model: model, Table: table,
		Columns:          iFields,
		ColumnsSql:       columnsSql,
		ColumnsInsertSql: columnsInsertSql,
		FieldsSql:        fieldsSql}
}
