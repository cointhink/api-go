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
	stateName := algorunInstance.Status
	newStateName := proto.Algorun_States_name[int32(newStatus)]
	log.Printf("algorun %s update to %s", algorunInstance.Id, newStateName)

	if stateName == newStateName {
		return nil
	}

	var updateState *string
	// state machine goes here
	if newStateName == proto.Algorun_States_name[int32(proto.Algorun_deleted)] {
		if stateName == proto.Algorun_States_name[int32(proto.Algorun_building)] ||
			stateName == proto.Algorun_States_name[int32(proto.Algorun_running)] ||
			stateName == proto.Algorun_States_name[int32(proto.Algorun_stopped)] {
			updateState = &newStateName
		} else {
			log.Printf("algorun %s in bad state (%s) for delete", algorunInstance.Id, algorunInstance.Status)
		}
	} else if newStateName == proto.Algorun_States_name[int32(proto.Algorun_building)] {
		log.Printf("algorun %s update to state building (initial state) makes no sense",
			algorunInstance.Id, newStateName)
	} else if newStateName == proto.Algorun_States_name[int32(proto.Algorun_stopped)] {
		updateState = &newStateName
	} else {
		log.Printf("algorun %s unhandled newState %s", algorunInstance.Id, newStateName)
	}

	if updateState != nil {
		_, err := db.D.Handle.Exec("update algoruns set status = $1 where id = $2",
			*updateState, algorunInstance.Id)
		if err != nil {
			log.Printf("%v", err)
		} else {
			log.Printf("algorun %s updated from %s to %s", algorunInstance.Id, stateName, *updateState)
			algorunInstance.Status = *updateState
		}
	}
	return nil
}
