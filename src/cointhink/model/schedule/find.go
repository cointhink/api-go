package schedule

import "cointhink/db"
import "cointhink/proto"
import "cointhink/constants"
import "log"
import "strconv"
import "time"

func Find(scheduleId string) (proto.Schedule, error) {
	schedule := proto.Schedule{}
	err := db.D.Handle.Get(&schedule,
		"select "+Columns+" from schedules where id = $1",
		scheduleId)
	if err != nil {
		log.Printf("schedule.Find SQL: %v", err)
		return schedule, err
	} else {
		return schedule, nil
	}
}

func FindWithAccount(scheduleId string, accountId string) (proto.Schedule, error) {
	schedule := proto.Schedule{}
	err := db.D.Handle.Get(&schedule,
		"select "+Columns+" from schedules where id = $1 and account_id = $2",
		scheduleId, accountId)
	if err != nil {
		log.Printf("schedule.FindWithAccount SQL: %v", err)
		return schedule, err
	} else {
		return schedule, nil
	}
}

func List(accountId string) ([]*proto.Schedule, error) {
	schedules := []*proto.Schedule{}
	err := db.D.Handle.Select(&schedules,
		"select "+Columns+" from schedules where status != "+
			strconv.FormatInt(int64(proto.Schedule_States_value["deleted"]), 10)+" and account_id = $1",
		accountId)
	if err != nil {
		log.Printf("ScheduleFind SQL: %v", err)
		return schedules, err
	} else {
		return schedules, nil
	}
}

func EnabledNow(_schedule *proto.Schedule) bool {
	enabledUntil, err := time.Parse(constants.ISO8601, _schedule.EnabledUntil)
	if err != nil {
		log.Printf("schedule.EnabledNow err: %v", err)
		return false
	} else {
		now := time.Now()
		log.Printf("schedule.EnabledNow now: %+v enabledUntil: %+v (%+v)", now, enabledUntil, _schedule.Id)
		return enabledUntil.After(now)
	}
}

func RunningExpireds(time time.Time) []*proto.Schedule {
	schedules := []*proto.Schedule{}
	err := db.D.Handle.Select(&schedules,
		"select "+Columns+" from schedules where status = "+
			strconv.FormatInt(int64(proto.Schedule_States_value["enabled"]), 10)+" and "+
			" enabled_until < $1", time)
	if err != nil {
		log.Printf("RunningExpireds SQL: %v", err)
	}
	return schedules
}
