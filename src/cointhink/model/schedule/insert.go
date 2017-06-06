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
		new_id, _ := result.LastInsertId()
		log.Printf("Schedule new id %s", new_id)
		log.Printf("Schedule Create result: %+v", result)
	}
	return nil
}
