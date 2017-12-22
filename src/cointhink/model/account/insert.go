package account

import "cointhink/proto"
import "cointhink/db"
import "log"

var Schema db.SqlDetail = db.Register(proto.Account{})

func Insert(item *proto.Account) error {
	item.Id = db.NewId(Schema.Table)
	sql := "insert into " + Schema.Table + " (" + Schema.ColumnsInsertSql + ") " +
		"values (" + Schema.FieldsSql + ")"
	_, err := db.D.Handle.NamedExec(sql, item)
	if err != nil {
		log.Printf(Schema.Table+" Create err: %v\n%s", err, sql)
		return err
	}
	return nil
}

func UpdateStripe(_account *proto.Account, newStripe string) error {
	sql := "update " + Schema.Table + " set stripe = $1 where id = $2"
	_, err := db.D.Handle.Exec(sql, newStripe, _account.Id)
	if err != nil {
		log.Printf("account.UpdateStripe err %v", err)
		return err
	}
	_account.Stripe = newStripe
	return nil
}
