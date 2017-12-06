package token

import "cointhink/db"
import "cointhink/proto"
import "log"

func FindByAccountId(accountId string, algorunId string) (*proto.Token, error) {
	log.Printf("token.FindByAccount accountId %+v algorunId %+v", accountId, algorunId)
	item := &proto.Token{}
	err := db.D.Handle.Get(item,
		"select "+Columns+" from "+Table+" where account_id = $1 and algorun_id = $2",
		accountId, algorunId)
	if err != nil {
		return nil, err
	} else {
		return item, nil
	}
}

func FindByToken(token_str string) (*proto.Token, error) {
	item := &proto.Token{}
	err := db.D.Handle.Get(item,
		"select "+Columns+" from "+Table+" where token = $1", token_str)
	if err != nil {
		return item, err
	} else {
		return item, nil
	}
}
