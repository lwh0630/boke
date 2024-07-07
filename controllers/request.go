package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID" // 为 gin.Context.Keys 指定一个键值

var ErrUserNotLogin = errors.New("用户未登录")

// GetCurrentUserID 获取当前登录的用户ID
func GetCurrentUserID(c *gin.Context) (uid int64, err error) {
	uidAny, ok := c.Get(CtxUserIDKey)
	if !ok {
		return -1, ErrUserNotLogin
	}
	uid, ok = uidAny.(int64)
	if !ok {
		return -1, ErrUserNotLogin
	}
	return uid, nil
}
