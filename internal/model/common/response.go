package common

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Result struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	RequestId string      `json:"request_id"`
}

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func NewResult() *Result {
	return &Result{
		Code:      20000,
		Message:   "SUCCESS",
		Data:      nil,
		RequestId: "",
	}
}

func (r *Response) Success() {
	result := NewResult()
	result.Code = 20000
	result.Message = "SUCCESS"
	result.RequestId = r.Ctx.GetString("request_id")
	r.Ctx.JSON(http.StatusOK, result)
}

func (r *Response) SuccessData(data interface{}) {
	result := NewResult()
	result.Code = 20000
	result.Message = "SUCCESS"
	result.Data = data
	result.RequestId = r.Ctx.GetString("request_id")
	log.WithField("request_id", result.RequestId).Info(result.Data)
	r.Ctx.JSON(http.StatusOK, result)
}

func (r *Response) Fail(message string) {
	result := NewResult()
	result.Code = 40000
	result.Message = message
	if result.Message == "" {
		result.Message = "fail"
	}
	result.RequestId = r.Ctx.GetString("request_id")
	log.WithField("request_id", result.RequestId).Error(result.Message)
	r.Ctx.JSON(http.StatusOK, result)
}

func (r *Response) Error(message error) {
	result := NewResult()
	result.Code = 40000
	result.Message = message.Error()
	result.RequestId = r.Ctx.GetString("request_id")
	log.WithField("request_id", result.RequestId).Error(result.Message)
	r.Ctx.JSON(http.StatusOK, result)
}

func (r *Response) SetResult(code int, message string, data interface{}) {
	var result = &Result{
		Code:      code,
		Message:   message,
		Data:      data,
		RequestId: r.Ctx.GetString("request_id"),
	}
	r.send(result)
}

func (r *Response) send(result *Result) {
	log.WithField("request_id", result.RequestId).Error(result.Message)
	r.Ctx.JSON(http.StatusOK, result)
}
