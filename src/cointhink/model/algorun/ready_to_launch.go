package algorun

import "log"
import "errors"

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
