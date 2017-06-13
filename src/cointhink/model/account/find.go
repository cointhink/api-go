package account

import "cointhink/db"
import "cointhink/proto"
import "log"

var Columns = "id, email, fullname"

func Find(accountId string) (proto.Account, error) {
	account := proto.Account{}
	err := db.D.Handle.Get(&account,
		"select "+Columns+" from accounts where id = $1",
		accountId)
	if err != nil {
		log.Printf("account.Find SQL: %v", err)
		return account, err
	} else {
		return account, nil
	}
}
