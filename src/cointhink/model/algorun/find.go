package algorun

import "cointhink/db"
import "cointhink/proto"
import "log"

func Find(id string) (*proto.Algorun, error) {
	run := &proto.Algorun{}
	err := db.D.Handle.Get(run,
		"select "+Columns+" from algoruns where id = $1", id)
	if err != nil {
		return run, err
	} else {
		return run, nil
	}
}

func FindFromSchedule(accountId string, scheduleId string) (*proto.Algorun, error) {
	run := &proto.Algorun{}
	err := db.D.Handle.Get(run,
		"select "+Columns+" from algoruns where account_id = $1 and schedule_id = $2",
		accountId, scheduleId)
	if err != nil {
		return run, err
	} else {
		return run, nil
	}
}

func FindReady(accountId string, scheduleId string) ([]*proto.Algorun, error) {
	ids := []*proto.Algorun{}
	err := db.D.Handle.Select(&ids,
		"select "+Columns+" from algoruns where status != 'deleted' and account_id = $1 and schedule_id = $2",
		accountId, scheduleId)
	if err != nil {
		return ids, err
	} else {
		return ids, nil
	}
}

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
