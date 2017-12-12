package actions

import "log"

import "cointhink/proto"

import gproto "github.com/golang/protobuf/proto"

func DoLambdaResponse(_lambda_response *proto.LambdaResponse, token *proto.Token) []gproto.Message {
	var responses []gproto.Message

	log.Printf("LambdaResponse %s ", _lambda_response.StateOut)

	return responses
}
