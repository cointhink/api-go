package model

import "cointhink/db"
import "errors"

func TokenFindAccountId(token string) (string, error) {
	rows, err := db.D.Handle.Query(
		"select account_id from tokens where token = $1",
		token)
	if err != nil {
		return "", err
	}
	if rows.Next() {
		var id string
		rows.Scan(&id)
		rows.Close()
		return id, nil
	} else {
		rows.Close()
		return "", errors.New("token not found")
	}
}
