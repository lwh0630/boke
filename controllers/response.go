package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	{
		"code": 1001	// 错误码
		"msg": xx, 		// 提示中的信息
		"data": {}		// 数据
	}
*/

type ResponseController struct {
	Code CodeRes     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseController{
		Code: CodeSuccess,
		Msg:  codeMsg[CodeSuccess],
		Data: data,
	})
}

func ResponseError(c *gin.Context, code CodeRes) {
	c.JSON(http.StatusBadRequest, &ResponseController{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code CodeRes, msg interface{}) {
	c.JSON(http.StatusBadRequest, &ResponseController{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
