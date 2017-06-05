package algorun

import "cointhink/db"
import "log"
import "errors"

func ReadyToLaunch(accountId string, algorithmId string) error {
	ids := []string{}
	err := db.D.Handle.Select(&ids,
		"select id from algoruns where account_id = $1 and algorithm_id = $2",
		accountId, algorithmId)
	if err != nil {
		return err
	} else {
		log.Printf("algoruns found %v", ids)
		if len(ids) > 0 {
			return errors.New("already running")
		}
	}
	return nil
}
