package account

import "cointhink/db"
import "cointhink/proto"
import "log"

func Find(accountId string) (proto.Account, error) {
	account := proto.Account{}
	sql := "select " + Schema.ColumnsSql + " from " + Schema.Table + " where id = $1"
	err := db.D.Handle.Get(&account, sql, accountId)
	if err != nil {
		log.Printf("account.Find SQL: %v", err)
		return account, err
	} else {
		return account, nil
	}
}

func FindByEmail(email string) (proto.Account, error) {
	account := proto.Account{}
	err := db.D.Handle.Get(&account,
		"select "+Schema.ColumnsSql+" from "+Schema.Table+" where email = $1",
		email)
	if err != nil {
		log.Printf("account.Find SQL: %v", err)
		return account, err
	} else {
		return account, nil
	}
}
