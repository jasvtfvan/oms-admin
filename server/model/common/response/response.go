package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// 通过了middleware
const (
	SUCCESS = http.StatusOK
	ERROR   = http.StatusBadRequest
)

// 未通过middleware
const (
	UN_AUTH = http.StatusUnauthorized
	BAD     = http.StatusBadRequest
)

/*
通过了middleware的result
*/
func result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}
func Success(data interface{}, message string, c *gin.Context) {
	result(SUCCESS, data, message+"!", c)
}
func Fail(data interface{}, message string, c *gin.Context) {
	result(ERROR, data, message+"!", c)
}

/*
没通过middleware的result
*/
func resultJwt(status int, data interface{}, msg string, c *gin.Context) {
	c.AbortWithStatusJSON(status, Response{
		status,
		data,
		msg,
	})
}
func Unauthorized(data interface{}, message string, c *gin.Context) {
	resultJwt(UN_AUTH, data, message+"!", c)
}
func BadReq(data interface{}, message string, c *gin.Context) {
	resultJwt(BAD, data, message+"!", c)
}
