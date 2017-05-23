package model

import "cointhink/proto"
import "cointhink/db"
import "errors"

func AccountFind(id string) (string, error) {
	rows, _ := db.D.Handle.Query(
		"select id from accounts where id = $1",
		id)
	if rows.Next() {
		var id string
		rows.Scan(&id)
		return id, nil
	} else {
		return "", errors.New("account id not found")
	}
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
		return "", errors.New("account email not found")
	}
}

func AccountSave(account proto.Account) {
}
