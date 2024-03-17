package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	HttpCode int         `json:"-"`
	Code     ResCode     `json:"code"`
	Message  any         `json:"message"`
	Data     interface{} `json:"data,omitempty"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		HttpCode: c.Writer.Status(),
		Code:     CodeSuccess,
		Message:  CodeSuccess.Msg(),
		Data:     data,
	})
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &Response{
		HttpCode: c.Writer.Status(),
		Code:     code,
		Message:  code.Msg(),
		Data:     nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg any) {
	c.JSON(http.StatusOK, &Response{
		HttpCode: c.Writer.Status(),
		Code:     code,
		Message:  msg,
		Data:     nil,
	})
}
