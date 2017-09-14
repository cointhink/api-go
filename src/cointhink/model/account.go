package model

import "cointhink/proto"
import "cointhink/db"
import "errors"

func AccountFindByEmail(email string) (string, error) {
	rows, err := db.D.Handle.Query(
		"select id from accounts where email = $1",
		email)
	if err != nil {
		return "", err
	} else {
		if rows.Next() {
			var id string
			rows.Scan(&id)
			return id, nil
		} else {
			return "", errors.New("account email not found")
		}
	}
}

func AccountSave(account proto.Account) {
}
