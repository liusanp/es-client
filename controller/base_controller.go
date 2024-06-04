package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (b *BaseController) Ok(c *gin.Context, msg string, data interface{}) {
	res := Response{
		Code: 0,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, res)
}

func (b *BaseController) Error(c *gin.Context, msg string) {
	res := Response{
		Code: 10001,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusBadRequest, res)
}

func (b *BaseController) Return(c *gin.Context, code int, msg string, data interface{}) {
	res := Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, res)
}
