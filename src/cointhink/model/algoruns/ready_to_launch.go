package algoruns

import "cointhink/db"
import "errors"

func ReadyToLaunch(id string) error {
	rows, _ := db.D.Handle.Query(
		"select id from algoruns where algorithm_id = $1",
		id)
	if rows.Next() {
		var id string
		rows.Scan(&id)
		return nil
	} else {
		return errors.New("algoruns: algorithm id not found")
	}
}
