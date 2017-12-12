package actions

import "log"

import "cointhink/proto"
import "cointhink/model/algorun"
import "cointhink/model/token"

import gproto "github.com/golang/protobuf/proto"

func DoLambdaResponse(_lambda_response *proto.LambdaResponse, _token *proto.Token) []gproto.Message {
	var responses []gproto.Message

	log.Printf("LambdaResponse %s %s ", _token.AlgorunId, _lambda_response.StateOut)
	_ltoken, err := token.FindByToken(_lambda_response.Token)
	if err != nil {
		log.Printf("dolambdaresponse token %#v err %#v", _lambda_response.Token, err)
	} else {
		_algorun, err := algorun.Find(_ltoken.AlgorunId)
		if err != nil {
			log.Printf("dolambdaresponse algorun err %#v", err)
		} else {
			algorun.UpdateState(_algorun, _lambda_response.StateOut)
		}
	}

	return responses
}
