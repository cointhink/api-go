package credit_journal

import "cointhink/db"
import "cointhink/proto"

func TotalCredits(account *proto.Account) (int32, error) {
	journals := []*proto.CreditJournal{}
	err := db.D.Handle.Select(journals,
		"select "+Columns+" from "+Table+" where id = $1", account.Id)
	if err != nil {
		return 0, err
	} else {
		return 0, nil
	}
}
