package algorun

import "cointhink/proto"
import "cointhink/db"
import "log"

var Columns = "id, algorithm_id, account_id, schedule_id, status"

func Insert(algorunInstance *proto.Algorun) error {
	algorunInstance.Id = db.NewId("algoruns")
	_, err := db.D.Handle.NamedExec("insert into algoruns ("+Columns+") "+
		"values (:id, :algorithm_id, :account_id, :schedule_id, :status)", algorunInstance)
	if err != nil {
		log.Printf("algorun Create err: %v", err)
		return err
	}
	return nil
}

func UpdateStatus(algorunInstance *proto.Algorun, newStatus proto.Algorun_States) error {
	log.Printf("UpdateStatus %s to %v", algorunInstance.Id, newStatus)
	// go type screwery
	currentStatusInt := proto.Algorun_States_value[algorunInstance.Status]
	newStatusInt := int32(newStatus)
	if currentStatusInt == newStatusInt {
		return nil
	}

	// state machine goes here
	if newStatus == proto.Algorun_deleted {
		if currentStatusInt == int32(proto.Algorun_running) ||
			currentStatusInt == int32(proto.Algorun_stopped) {
			newStateName := proto.Algorun_States_name[newStatusInt]
			row, err := db.D.Handle.Exec("update algoruns set status = $1 where id = $2",
				newStateName, algorunInstance.Id)
			if err != nil {
				log.Printf("%v", err)
			} else {
				log.Printf("%v", row)
			}
		} else {
			log.Printf("algorun %s in bad state (%s) for delete", algorunInstance.Id, algorunInstance.Status)
		}
	}
	return nil
}
