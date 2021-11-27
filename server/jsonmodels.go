package server

type ErrorCode = int

const (
	EOK ErrorCode = iota

	EBadParse
	EBadSerialization
	EBadOptomization
	EGeneric
)

type ResponseBody struct {
	Body         interface{} `json:"message_body" bson:"message_body,inline,omitempty"`
	Code         ErrorCode   `json:"error_code" bson:"error_code"`
	ErrorMessage string      `json:"error_message" bson:"error_message`
}
