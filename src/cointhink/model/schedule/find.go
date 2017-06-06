package schedule

import "cointhink/db"
import "cointhink/proto"
import "log"

func Find(scheduleId string, accountId string) (proto.Schedule, error) {
	schedule := proto.Schedule{}
	err := db.D.Handle.Get(&schedule,
		"select * from schedules where id = $1 and account_id = $2",
		scheduleId, accountId)
	if err != nil {
		log.Printf("ScheduleFind SQL: %v", err)
		return schedule, err
	} else {
		return schedule, nil
	}
}

func List(accountId string) ([]*proto.Schedule, error) {
	schedules := []*proto.Schedule{}
	err := db.D.Handle.Select(&schedules,
		"select id, account_id, algorithm_id, status, initial_state from schedules where account_id = $1",
		accountId)
	if err != nil {
		log.Printf("ScheduleFind SQL: %v", err)
		return schedules, err
	} else {
		return schedules, nil
	}
}
