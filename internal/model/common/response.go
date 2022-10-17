package common

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	RequestId string      `json:"request_id"`
}

type Result struct {
	Ctx *gin.Context
}

func NewResult(ctx *gin.Context) *Result {
	return &Result{Ctx: ctx}
}

func NewResponse() *Response {
	return &Response{
		Code:      20000,
		Message:   "SUCCESS",
		Data:      nil,
		RequestId: "",
	}
}

func (r *Result) Success() {
	response := NewResponse()
	response.Code = 20000
	response.Message = "SUCCESS"
	response.RequestId = r.Ctx.GetString("request_id")
	r.Ctx.JSON(http.StatusOK, response)
}

func (r *Result) SuccessData(data interface{}) {
	response := NewResponse()
	response.Code = 20000
	response.Message = "SUCCESS"
	response.Data = data
	response.RequestId = r.Ctx.GetString("request_id")
	log.WithField("request_id", response.RequestId).Info(response.Data)
	r.Ctx.JSON(http.StatusOK, response)
}

func (r *Result) Fail(message string) {
	response := NewResponse()
	response.Code = 40000
	response.Message = message
	if response.Message == "" {
		response.Message = "fail"
	}
	response.RequestId = r.Ctx.GetString("request_id")
	log.WithField("request_id", response.RequestId).Error(response.Message)
	r.Ctx.JSON(http.StatusOK, response)
}

func (r *Result) Error(message error) {
	response := NewResponse()
	response.Code = 40000
	response.Message = message.Error()
	response.RequestId = r.Ctx.GetString("request_id")
	log.WithField("request_id", response.RequestId).Error(response.Message)
	r.Ctx.JSON(http.StatusOK, response)
}

func (r *Result) SetResult(code int, message string, data interface{}) {
	var response = &Response{
		Code:      code,
		Message:   message,
		Data:      data,
		RequestId: r.Ctx.GetString("request_id"),
	}
	r.send(response)
}

func (r *Result) send(response *Response) {
	log.WithField("request_id", response.RequestId).Error(response.Message)
	r.Ctx.JSON(http.StatusOK, response)
}
