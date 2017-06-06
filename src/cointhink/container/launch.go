package container

import "cointhink/net"
import "cointhink/model/algorun"
import "cointhink/proto"
import "errors"
import "log"

func Launch(accountId string, algorithmId string) error {
	err := algorun.ReadyToLaunch(accountId, algorithmId)
	if err != nil {
		log.Printf("Launch: algo not ready. %v", err)
		return errors.New("Launch: algo not ready")
	} else {
		log.Printf("Launch: algo ready. launching")
		_algorun := proto.Algorun{AccountId: accountId, AlgorithmId: algorithmId}
		algorun.Insert(&_algorun)
		net.LxdLaunch(net.Lxc{Name: _algorun.Id,
			Source: net.LxcSource{Type: "image", Fingerprint: "6978077ac9f8"}})
	}
	return nil
}
