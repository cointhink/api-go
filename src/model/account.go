package model

import "proto"
import "db"
import "errors"

func AccountSave(account proto.Account) {
}

func AccountFindByEmail(email string) (string, error) {
	rows, _ := db.D.Handle.Query(
		"select id from accounts where email = $1",
		email)
	if rows.Next() {
		var id string
		rows.Scan(&id)
		return id, nil
	} else {
		return "", errors.New("account id not found")
	}
}
