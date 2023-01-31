package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code    uint        `json:"code"`    // 状态码
	Message string      `json:"message"` // 消息
	Data    interface{} `json:"data"`    // 返回数据
}

const (
	SUCCESS uint = 200
	ERROR   uint = 400
)

func H(result Result) map[string]interface{} {
	return gin.H{"code": result.Code, "message": result.Message, "data": result.Data}
}

func Success(message string, data interface{}) Result {
	var result Result
	result.Code = SUCCESS
	result.Message = message
	result.Data = data
	return result
}

func SuccessNoData(message string) Result {
	var result Result
	result.Code = SUCCESS
	result.Message = message
	result.Data = nil
	return result
}

func Error(message string) Result {
	var result Result
	result.Code = ERROR
	result.Message = message
	result.Data = nil
	return result
}

func DealResult(c *gin.Context, result Result) {
	c.JSON(http.StatusOK, H(result))
}
