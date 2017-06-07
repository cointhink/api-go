package container

import "log"

import "cointhink/lxd"

func Stop(algorunId string) error {
	log.Printf("Stop: %s", algorunId)
	lxd.Delete(algorunId)
	return nil
}
