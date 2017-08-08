package actions

import "log"
import "cointhink/proto"
import gproto "github.com/golang/protobuf/proto"
import "cointhink/model/algorithm"

func DoAlgorithmList(_algorithmList *proto.AlgorithmList, accountId string) []gproto.Message {
	var responses []gproto.Message

	rows, err := algorithm.FindAll(accountId)
	if err != nil {
		log.Printf("algo err %+v", err)
		responses = append(responses, &proto.AlgorithmListResponse{Ok: false,
			Algorithms: []*proto.Algorithm{}})
	} else {
		log.Printf("algo rows %d", len(rows))

		responses = append(responses, &proto.AlgorithmListResponse{Ok: true,
			Algorithms: rows})
	}
	return responses
}
