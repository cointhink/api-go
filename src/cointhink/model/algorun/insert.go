package algorun

import "cointhink/proto"
import "cointhink/db"
import "log"

var Columns = "id, algorithm_id, account_id"

func Insert(algorunInstance *proto.Algorun) error {
	algorunInstance.Id = db.NewId("algoruns")
	result, err := db.D.Handle.NamedExec("insert into algoruns ("+Columns+") "+
		"values (:Id, :AlgorithmId, :AccountId)", algorunInstance)
	if err != nil {
		log.Printf("algorun Create err: %v", err)
		return err
	} else {
		log.Printf("algorun Create result: %+v", result)
	}
	return nil
}
