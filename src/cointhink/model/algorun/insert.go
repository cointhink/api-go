package algorun

import "cointhink/proto"
import "cointhink/db"
import "log"

var Columns = "id, algorithm_id, account_id, schedule_id, status, code, image"
var Fields = ":id, :algorithm_id, :account_id, :schedule_id, :status, :code, :image"
var Table = "algoruns"

func Insert(item *proto.Algorun) error {
	item.Id = db.NewId(Table)
	_, err := db.D.Handle.NamedExec("insert into "+Table+" ("+Columns+") "+"values ("+Fields+")", item)
	if err != nil {
		log.Printf("algorun Create err: %v", err)
		return err
	}
	return nil
}

func UpdateStatus(algorunInstance *proto.Algorun, newStatus proto.Algorun_States) error {
	stateName := algorunInstance.Status
	newStateName := proto.Algorun_States_name[int32(newStatus)]

	if stateName == newStateName {
		return nil
	}

	var updateState *string
	// state machine goes here
	if newStateName == proto.Algorun_States_name[int32(proto.Algorun_deleted)] {
		if stateName == proto.Algorun_States_name[int32(proto.Algorun_building)] ||
			stateName == proto.Algorun_States_name[int32(proto.Algorun_running)] ||
			stateName == proto.Algorun_States_name[int32(proto.Algorun_destroying)] ||
			stateName == proto.Algorun_States_name[int32(proto.Algorun_stopped)] {
			updateState = &newStateName
		} else {
			log.Printf("algorun %s in bad state (%s) for delete", algorunInstance.Id, algorunInstance.Status)
		}
	} else if newStateName == proto.Algorun_States_name[int32(proto.Algorun_building)] {
		log.Printf("algorun %s update to state building (initial state) makes no sense",
			algorunInstance.Id, newStateName)
	} else if newStateName == proto.Algorun_States_name[int32(proto.Algorun_starting)] {
		if stateName == proto.Algorun_States_name[int32(proto.Algorun_building)] {
			updateState = &newStateName
		} else {
			log.Printf("algorun %s update to starting from building (initial state) makes no sense",
				algorunInstance.Id)
		}
	} else if newStateName == proto.Algorun_States_name[int32(proto.Algorun_running)] {
		if stateName == proto.Algorun_States_name[int32(proto.Algorun_starting)] {
			updateState = &newStateName
		} else {
			log.Printf("algorun %s update to running from %s makes no sense",
				algorunInstance.Id, stateName)
		}
	} else if newStateName == proto.Algorun_States_name[int32(proto.Algorun_stopped)] {
		updateState = &newStateName
	} else if newStateName == proto.Algorun_States_name[int32(proto.Algorun_destroying)] {
		updateState = &newStateName
	} else {
		log.Printf("!algorun %s unhandled newState %s", algorunInstance.Id, newStateName)
	}

	if updateState != nil {
		_, err := db.D.Handle.Exec("update "+Table+" set status = $1 where id = $2",
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
