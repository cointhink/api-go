package algorithm

import "cointhink/proto"
import "cointhink/db"
import "log"

var Columns = "id, account_id, code, status, description"
var Fields = ":id, :account_id, :code, :status, :description"
var Table = "algorithms"

func Insert(item *proto.Algorithm) error {
	item.Id = db.NewId(Table)
	_, err := db.D.Handle.NamedExec("insert into "+Table+" ("+Columns+") "+"values ("+Fields+")", item)
	if err != nil {
		log.Printf(Table+" Create err: %v", err)
		return err
	}
	return nil
}
