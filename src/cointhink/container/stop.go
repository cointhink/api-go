package container

import "log"

import "cointhink/lxd"

func Stop(algorunId string) error {
	log.Printf("Stop: %s", algorunId)
	op := lxd.Delete(algorunId)
	if op.Type != "error" {
		lxd.LXDOPq <- op
	}
	return nil
}
