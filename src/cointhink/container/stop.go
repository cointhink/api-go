package container

import "log"

import "cointhink/lxd"
import "cointhink/proto"
import "cointhink/model/account"

func Stop(algorun *proto.Algorun) error {
	log.Printf("Stop: %s", algorun.Id)
	op := lxd.Delete(algorun.Id)
	if op.Type != "error" {
		_account, err := account.Find(algorun.AccountId)
		if err == nil {
			lxd.LXDOPq <- lxd.AccountOperation{Account: &_account, Operation: op}
		}
	}
	return nil
}
