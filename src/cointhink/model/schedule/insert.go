package schedule

import "cointhink/proto"
import "cointhink/db"
import "log"

var Columns = "id, account_id, algorithm_id, status, initial_state"

func Insert(scheduleInstance *proto.Schedule) error {
	scheduleInstance.Id = db.NewId("schedules")
	result, err := db.D.Handle.NamedExec("insert into schedules ("+Columns+") "+
		"values (:id, :account_id, :algorithm_id, :status, :initial_state)", scheduleInstance)
	if err != nil {
		log.Printf("Schedule Create err: %v", err)
		return err
	} else {
		log.Printf("Schedule Create result: %+v", result)
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
