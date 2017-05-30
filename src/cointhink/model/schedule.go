package model

import "cointhink/db"
import "cointhink/proto"
import "errors"
import "log"
import "database/sql"

func ScheduleFind(id string) (string, error) {
	rows, _ := db.D.Handle.Query(
		"select id from schedules where id = $1",
		id)
	if rows.Next() {
		var id string
		rows.Scan(&id)
		return id, nil
	} else {
		return "", errors.New("schedule id not found")
	}
}

func ScheduleList(accountId string) ([]*proto.Schedule, error) {
	rows, err := db.D.Handle.Query(
		"select id, account_id, algorithm_id, status from schedules where account_id = $1",
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

func SchedulePopulate(rows *sql.Rows, schedule *proto.Schedule) error {
	return rows.Scan(
		schedule.Id,
		schedule.AccountId,
		schedule.AlgorithmId,
		schedule.Status,
	)
}

func ScheduleInsert(accountId string, algorithmId string, status string) error {
	stmt, err := db.D.Handle.Prepare(
		"insert into schedules (id, account_id, algorithm_id, status) values ($1, $2, $3, $4)")
	if err != nil {
		log.Printf("prepare %+v", err)
		return err
	}

	id := db.NewId("schedules")
	sql_result, err := stmt.Exec(
		id,
		accountId,
		algorithmId,
		status)
	stmt.Close()
	if err != nil {
		log.Printf("upsert %+v", err)
		return err
	}
	new_id, err := sql_result.LastInsertId()
	log.Printf("Account new id %s", new_id)
	return nil
}
