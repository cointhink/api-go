package algolog

import "cointhink/db"
import "cointhink/proto"

func Find(id string) (*proto.Algorun, error) {
	run := &proto.Algorun{}
	err := db.D.Handle.Get(run,
		"select "+Columns+" from algologs where id = $1", id)
	if err != nil {
		return run, err
	} else {
		return run, nil
	}
}
