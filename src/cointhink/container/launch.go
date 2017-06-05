package container

import "cointhink/net"
import "cointhink/model/algoruns"
import "errors"

func Launch(algorithmId string) error {
	err := algoruns.ReadyToLaunch(algorithmId)
	if err != nil {
		net.LxdLaunch(net.Lxc{Name: algorithmId,
			Source: net.LxcSource{Type: "image", Fingerprint: "6978077ac9f8"}})
	} else {
		return errors.New("not ready")
	}
	return nil
}
