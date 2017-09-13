package account

import "cointhink/proto"
import "cointhink/db"
import "log"

var Columns = "id, fullname, email, time_zone, schedule_credits"
var Fields = ":id, :fullname, :email, :time_zone, :schedule_credits"
var Table = "accounts"

func Insert(item *proto.Account) error {
	item.Id = db.NewId(Table)
	_, err := db.D.Handle.NamedExec("insert into "+Table+" ("+Columns+") "+"values ("+Fields+")", item)
	if err != nil {
		log.Printf(Table+" Create err: %v", err)
		return err
	}
	return nil
}
