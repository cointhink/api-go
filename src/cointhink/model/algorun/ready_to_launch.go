package algorun

import "cointhink/db"
import "log"
import "errors"

func Find(accountId string, scheduleId string) ([]string, error) {
	ids := []string{}
	err := db.D.Handle.Select(&ids,
		"select id from algoruns where account_id = $1 and schedule_id = $2",
		accountId, scheduleId)
	if err != nil {
		return ids, err
	} else {
		return ids, nil
	}
}

func FindReady(accountId string, scheduleId string) ([]string, error) {
	ids := []string{}
	err := db.D.Handle.Select(&ids,
		"select id from algoruns where status != 'deleted' and account_id = $1 and schedule_id = $2",
		accountId, scheduleId)
	if err != nil {
		return ids, err
	} else {
		return ids, nil
	}
}

func ReadyToLaunch(accountId string, scheduleId string) error {
	ids, err := FindReady(accountId, scheduleId)
	if err != nil {
		return err
	} else {
		log.Printf("ReadytoLaunch found %v", ids)
		if len(ids) > 0 {
			return errors.New("already running")
		}
	}
	return nil
}
