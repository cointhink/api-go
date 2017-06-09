package container

import "log"
import "cointhink/model/algorun"
import "cointhink/lxd"
import "cointhink/proto"

func SyncAll() {
	runs, err := algorun.List()
	if err != nil {
	}

	log.Printf("Syncing %d algoruns", len(runs))
	for _, run := range runs {
		Sync(run)
	}
}

func Sync(run *proto.Algorun) {
	response, err := lxd.Status(run.Id)
	if err != nil {
		log.Printf("%v %v", run.Id, err)
	} else {
		log.Printf("C: %v %s %s", run.Id, run.Status, response.Status)
		if response.StatusCode == 404 {
			algorun.UpdateStatus(run, proto.Algorun_deleted)
		}
	}
}
