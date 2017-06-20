package container

import "log"

import "cointhink/lxd"
import "cointhink/q"
import "cointhink/proto"
import "cointhink/model/algorun"

func Stop(_algorun *proto.Algorun) error {
	log.Printf("Stop: %s", _algorun.Id)
	algorun.UpdateStatus(_algorun, proto.Algorun_destroying)
	op := lxd.Stop(_algorun.Id)
	if op.Type != "error" {
		q.LXDOPq <- q.AccountOperation{Algorun: _algorun, Operation: op}
	}
	return nil
}
