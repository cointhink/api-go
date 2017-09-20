package token

import "cointhink/proto"
import "cointhink/db"
import "log"
import "github.com/satori/go.uuid"

var Columns = "token, account_id, algorun_id"
var Fields = ":token, :account_id, :algorun_id"
var Table = "tokens"

func Insert(item *proto.Token) error {
	item.Id = db.NewId(Table)
	item.Token = uuid.NewV4().String()
	_, err := db.D.Handle.NamedExec("insert into "+Table+" ("+Columns+") "+"values ("+Fields+")", item)
	if err != nil {
		log.Printf(Table+" Create err: %v", err)
		return err
	} else {
		return nil
	}
}
