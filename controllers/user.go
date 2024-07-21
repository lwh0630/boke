package controllers

import (
	"bluebell/dataset/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// RegisterHandler /*
// 处理注册请求
func RegisterHandler(c *gin.Context) {
	// 1.参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求错误
		zap.L().Error("处理注册请求,请求参数无效", zap.Error(err))
		// 判断错误类型是否为validator类型
		if errs, ok := err.(validator.ValidationErrors); !ok {
			ResponseError(c, CodeInvalidParams)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		}
		return
	}
	// 2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("注册失败", zap.Error(err))
		if errors.Is(err, mysql.ErrUserExist) {
			ResponseError(c, CodeUserExist)
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	// 3.返回响应
	ResponseOk(c, nil)
}

func LoginHandler(c *gin.Context) {
	// 1.获取请求参数以及参数校验
	p := new(models.ParamLogin)
	if er := c.ShouldBindJSON(p); er != nil {
		zap.L().Error("处理错误，请求参数无效", zap.Error(er))
		// 判断错误类型是否为validator类型
		if err, ok := er.(validator.ValidationErrors); !ok {
			ResponseError(c, CodeInvalidParams)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(err.Translate(trans)))
		}
		return
	}
	// 2.处理逻辑
	user, err := logic.Login(p)
	if err != nil {
		if errors.Is(err, mysql.ErrUserNotExist) {
			ResponseError(c, CodeUserNotExist)
		} else if errors.Is(err, mysql.ErrInvalidPassWord) {
			ResponseError(c, CodeErrorPassword)
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	// 3.返回响应
	ResponseOk(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
}
