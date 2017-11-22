package schedule

import "cointhink/proto"
import "cointhink/db"
import "cointhink/constants"
import "cointhink/model/credit_journal"
import "log"
import "time"
import "errors"

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

func UpdateStatus(_schedule *proto.Schedule, newState proto.Schedule_States) {
	log.Printf("schedule.UpdateStatus %s to %v", _schedule.Id, newState)
	_, err := db.D.Handle.Exec("update schedules set status = $1 where id = $2",
		newState, _schedule.Id)
	if err != nil {
		log.Printf("schedule.UpdateState err %v", err)
	}
}

func UpdateInitialState(_schedule *proto.Schedule, initialState string) {
	log.Printf("schedule.UpdateInitialState %s to %v", _schedule.Id, initialState)
	_, err := db.D.Handle.Exec("update schedules set initial_state = $1 where id = $2",
		initialState, _schedule.Id)
	if err != nil {
		log.Printf("schedule.UpdateState err %v", err)
	}
}

func UpdateEnabledUntil(_schedule *proto.Schedule, until time.Time) {
	timeStr := until.UTC().Format(constants.ISO8601)
	log.Printf("schedule.UpdateStatus %s to %v", _schedule.Id, timeStr)
	_, err := db.D.Handle.Exec("update schedules set enabled_until = $1 where id = $2",
		timeStr, _schedule.Id)
	if err != nil {
		log.Printf("schedule.UpdateState err %v", err)
	}
}

func EnableUntilNextMonth(_schedule *proto.Schedule, _account *proto.Account) error {
	if _account.ScheduleCredits > 0 {
		c_err := credit_journal.Debit(_account, _schedule, 1)
		if c_err != nil {
			log.Printf("DoScheduleCreate credit_journal Debit err %+v", c_err)
		} else {
			future := time.Now().Add(24 * 30 * time.Hour)
			UpdateEnabledUntil(_schedule, future)
		}
	} else {
		return errors.New("Insufficient credits")
	}
	return nil
}
