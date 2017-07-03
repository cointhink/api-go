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
