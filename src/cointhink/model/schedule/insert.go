package schedule

import "cointhink/proto"
import "cointhink/db"
import "log"

var Columns = "id, account_id, algorithm_id, status, initial_state, enabled_until"
var Fields = ":id, :account_id, :algorithm_id, :status, :initial_state, :enabled_until"
var Table = "schedules"

func Insert(item *proto.Schedule) error {
	item.Id = db.NewId(Table)
	_, err := db.D.Handle.NamedExec("insert into "+Table+" ("+Columns+") "+"values ("+Fields+")", item)
	if err != nil {
		log.Printf(Table+" Create err: %v", err)
		return err
	}
	return nil
}

func UpdateStatus(_schedule proto.Schedule, newState proto.Schedule_States) {
	log.Printf("schedule.UpdateStatus %s to %v", _schedule.Id, newState)
	_, err := db.D.Handle.Exec("update schedules set status = $1 where id = $2",
		newState, _schedule.Id)
	if err != nil {
		log.Printf("schedule.UpdateState err %v", err)
	}
}
