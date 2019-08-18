package util

import (
	"biligo/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HTTP 统一返回值
type Result struct {
	Code      int32       `json:"code"`
	Message   string      `json:"message"`
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

// 构造一个成功返回值
func SuccessResult(data interface{}) *Result {
	return &Result{constant.ResultCodeSuccess,
		"处理成功",
		true,
		data,
		CurrentTimestamp()}
}

// 构造一个成功返回值 带msg
func SuccessResultWithMessage(msg string, data interface{}) *Result {
	return &Result{constant.ResultCodeSuccess,
		msg,
		true,
		data,
		CurrentTimestamp()}
}

// 构造一个错误返回值
func FailResult(data interface{}) *Result {
	return &Result{constant.ResultCodeFail,
		"处理失败",
		false,
		data,
		CurrentTimestamp()}
}

// 构造一个错误返回值 带msg
func FailResultWithMessage(msg string, data interface{}) *Result {
	return &Result{constant.ResultCodeFail,
		msg,
		false,
		data,
		CurrentTimestamp()}
}

// 构造一个错误返回值 自定义code 和 msg
func FailResultWithCodeAndMessage(code int32, msg string, data interface{}) *Result {
	return &Result{code,
		msg,
		false,
		data,
		CurrentTimestamp()}
}

// 返回 HTTP 响应
func (result *Result) ToJSON(c *gin.Context) {
	c.JSON(http.StatusOK, result)
}

// 返回 HTTP 响应及状态码
func (result *Result) ToJSONWithHttpStatus(c *gin.Context) {
	c.JSON(int(result.Code), result)
}
