package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Response struct {
	Code     Code        `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	DateTime string      `json:"date_time"`
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, newResponse(SuccessCode, data))
}

func Fail(ctx *gin.Context, code Code, data interface{}) {
	ctx.JSON(http.StatusBadRequest, newResponse(code, data))
}

func newResponse(code Code, data interface{}) *Response {
	return &Response{
		Code:     code,
		Message:  code.Message(),
		Data:     data,
		DateTime: time.Now().Format(time.RFC3339),
	}
}
