package container

import "log"

import "cointhink/lxd"
import "cointhink/proto"

func Stop(algorun *proto.Algorun) error {
	log.Printf("Stop: %s", algorun.Id)
	//op := lxd.Stop(algorun.Id) // delete uses force
	op := lxd.Delete(algorun.Id)
	if op.Type != "error" {
		lxd.LXDOPq <- lxd.AccountOperation{Algorun: algorun, Operation: op}
	}
	return nil
}
