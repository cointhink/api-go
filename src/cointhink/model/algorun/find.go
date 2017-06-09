package algorun

import "cointhink/db"
import "cointhink/proto"
import "log"

func List() ([]*proto.Algorun, error) {
	items := []*proto.Algorun{}
	err := db.D.Handle.Select(&items,
		"select "+Columns+" from algoruns")
	if err != nil {
		log.Printf("ScheduleFind SQL: %v", err)
		return items, err
	} else {
		return items, nil
	}
}
