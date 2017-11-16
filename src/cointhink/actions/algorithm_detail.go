package actions

import "log"
import "cointhink/proto"
import gproto "github.com/golang/protobuf/proto"
import "cointhink/model/algorithm"

func DoAlgorithmDetail(_algorithmDetail *proto.AlgorithmDetail, accountId string) []gproto.Message {
	var responses []gproto.Message

	item, err := algorithm.Find(_algorithmDetail.AlgorithmId)
	if err != nil {
		log.Printf("algo err %+v", err)
		responses = append(responses, &proto.AlgorithmDetailResponse{Ok: false, Message: "Not Found"})
	} else {
		if item.AccountId == accountId {
			responses = append(responses, &proto.AlgorithmDetailResponse{Ok: true,
				Algorithm: item})
		} else {
			responses = append(responses, &proto.AlgorithmDetailResponse{Ok: false, Message: "No permission"})
		}
	}
	return responses
}
