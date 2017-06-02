package model

import "cointhink/db"
import "cointhink/proto"
import "errors"
import "log"
import "database/sql"

var ScheduleColumns = "id, account_id, algorithm_id, status, initial_state"

func SchedulePopulate(rows *sql.Rows, schedule *proto.Schedule) error {
	return rows.Scan(
		&schedule.Id,
		&schedule.AccountId,
		&schedule.AlgorithmId,
		&schedule.Status,
		&schedule.InitialState,
	)
}

func ScheduleFind(scheduleId string, accountId string) (proto.Schedule, error) {
	schedule := proto.Schedule{}
	rows, err := db.D.Handle.Query(
		"select "+ScheduleColumns+" from schedules where id = $1 and account_id = $2",
		scheduleId, accountId)
	if err != nil {
		log.Printf("ScheduleFind SQL: %v", err)
		return schedule, err
	} else {
		if rows.Next() {
			SchedulePopulate(rows, &schedule)
			return schedule, nil
		} else {
			return schedule, errors.New("schedule id not found")
		}
	}
}

func ScheduleList(accountId string) ([]*proto.Schedule, error) {
	rows, err := db.D.Handle.Query(
		"select id, account_id, algorithm_id, status, initial_state from schedules where account_id = $1",
		accountId)
	if err != nil {
		return nil, err
	}
	schedules := []*proto.Schedule{}
	for rows.Next() {
		newSchedule := proto.Schedule{}
		err = SchedulePopulate(rows, &newSchedule)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%v", newSchedule)
		schedules = append(schedules, &newSchedule)
	}
	rows.Close()
	return schedules, nil
}

func ScheduleInsert(accountId string,
	algorithmId string,
	status string,
	initialState string) error {
	stmt, err := db.D.Handle.Prepare(
		"insert into schedules " +
			"(" + ScheduleColumns + ") values " +
			"($1, $2, $3, $4, $5)")
	if err != nil {
		log.Printf("prepare %+v", err)
		return err
	}

	id := db.NewId("schedules")
	sql_result, err := stmt.Exec(
		id,
		accountId,
		algorithmId,
		status,
		initialState)
	stmt.Close()
	if err != nil {
		log.Printf("upsert %+v", err)
		return err
	}
	new_id, err := sql_result.LastInsertId()
	log.Printf("Account new id %s", new_id)
	return nil
}
