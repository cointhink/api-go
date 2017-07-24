package token

import "cointhink/db"
import "log"
import "github.com/satori/go.uuid"
import "errors"

func InsertToken(accountId string) string {
	stmt, err := db.D.Handle.Prepare("insert into tokens (token, account_id) values ($1, $2)")
	if err != nil {
		log.Printf("prepare %+v", err)
	}

	uuid_str := uuid.NewV4().String()
	_, err = stmt.Exec(uuid_str, accountId)

	return uuid_str
}

func Find(account_id string) (string, error) {
	rows, _ := db.D.Handle.Query(
		"select token from tokens where account_id = $1",
		account_id)
	if rows.Next() {
		var token string
		rows.Scan(&token)
		return token, nil
	} else {
		return "", errors.New("token for account id not found")
	}
}
