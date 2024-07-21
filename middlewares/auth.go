package middlewares

import (
	"bluebell/controllers"
	"bluebell/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式
		// 1. 放在请求头中
		// 2. 放在请求体中
		// 3. 放在URL中
		// 本函数的请求中放在GET请求头中，形式如: Authorization: Bearer xxx.xxx.xxx
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controllers.ResponseError(c, controllers.CodeNeedLogin)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		tokenString := parts[1]
		mc, err := jwt.ParseToken(tokenString)
		if err != nil {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		c.Set(controllers.CtxUserIDKey, mc.UserID)
		c.Next()
	}
}
