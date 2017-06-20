package q

import "cointhink/proto"

type AccountOperation struct {
	Algorun   *proto.Algorun
	Operation *OperationResponse
}
