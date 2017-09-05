package algorithm

import "cointhink/db"
import "cointhink/proto"

func Find(id string) (*proto.Algorithm, error) {
	run := &proto.Algorithm{}
	err := db.D.Handle.Get(run,
		"select "+Columns+" from "+Table+" where id = $1", id)
	if err != nil {
		return run, err
	} else {
		return run, nil
	}
}

func FindAll(accountId string) ([]*proto.Algorithm, error) {
	rows := []*proto.Algorithm{}
	err := db.D.Handle.Select(&rows,
		"select "+Columns+" from "+Table+" where account_id = $1 or true", accountId)
	if err != nil {
		return rows, err
	} else {
		return rows, nil
	}
}

func FindReady() ([]*proto.Algorithm, error) {
	rows := []*proto.Algorithm{}
	err := db.D.Handle.Select(&rows,
		"select "+Columns+" from "+Table+" where status = 'ready'")
	if err != nil {
		return rows, err
	} else {
		return rows, nil
	}
}
