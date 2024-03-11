package response

import "fmt"

type Code int // 状态码

var (
	SuccessCode       Code = 0
	InputInvalidError Code = 8020
	ShouldBindError   Code = 8021
	GetUserError      Code = 8022
	JsonMarshalError  Code = 8022
	CreateUserErr     Code = 8023
)

var codeMsgDict = map[Code]string{
	SuccessCode:       "ok",
	InputInvalidError: "input invalid",
	ShouldBindError:   "should bind failed",
	JsonMarshalError:  "json marshal failed",
	CreateUserErr:     "create user failed",
	GetUserError:      "get user failed",
}

func (c *Code) Message() string {
	if msg, ok := codeMsgDict[*c]; ok {
		return msg
	}
	return fmt.Sprintf("unknown error code %d", *c)
}

type ErrorCode struct {
	Code Code `json:"code"`
}

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Code.Message())
}
