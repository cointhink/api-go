package algolog

import "cointhink/db"
import "cointhink/proto"

func Find(id string) (*proto.Algolog, error) {
	run := &proto.Algolog{}
	err := db.D.Handle.Get(run,
		"select "+Columns+", created_at from algologs where id = $1", id)
	if err != nil {
		return run, err
	} else {
		return run, nil
	}
}

func FindAll(algorunId string, limit int) ([]*proto.Algolog, error) {
	rows := []*proto.Algolog{}
	err := db.D.Handle.Select(&rows,
		"select "+Columns+", created_at from algologs where "+
			"algorun_id = $1 order by created_at desc limit $2", algorunId, limit)
	if err != nil {
		return rows, err
	} else {
		return rows, nil
	}
}
