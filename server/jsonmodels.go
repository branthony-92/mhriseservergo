package server

import (
	"go.mongodb.org/mongo-driver/bson"
)

type BSONEncoder interface {
	Encode() *bson.D
}

type BSONDecoder interface {
	Decode(b *bson.D)
}

type ErrorCode = int

const (
	EGeneric  ErrorCode = -1
	EBadParse ErrorCode = iota
	EBadSerialization
	EBadOptomization
)

type ResponseBody struct {
	Body         BSONEncoder
	Code         ErrorCode
	ErrorMessage string
}

func (resp *ResponseBody) Encode() *bson.D {
	doc := bson.D{
		{"error_code", resp.Code},
		{"error_message", resp.ErrorMessage},
		{"message_body", resp.Body.Encode()},
	}
	return &doc
}
