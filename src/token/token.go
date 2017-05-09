package token

import "db"
import "log"
import "github.com/satori/go.uuid"

func MakeToken(id string) string {
	stmt, err := db.D.Handle.Prepare("insert into tokens (token, account_id) values ($1, $2)")
	if err != nil {
		log.Printf("prepare %+v", err)
	}

	uuid_str := uuid.NewV4().String()
	_, err = stmt.Exec(uuid_str, id)

	return uuid_str
}
