package routes

import (
	"bluebell/config"
	"bluebell/controllers"
	"bluebell/logger"
	"bluebell/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetupRoutes() (engine *gin.Engine, err error) {
	if strings.ToLower(config.Cfg.Mode) == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	// 初始化本地化模块
	err = controllers.InitTrans("zh")
	if err != nil {
		return nil, err
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 静态文件
	r.GET("/", func(c *gin.Context) {
		c.File("html/index.html")
	})
	// 设置静态文件目录
	r.Static("/static", "html/static")

	// 设置路由组
	v1 := r.Group("/api/v1")
	// 注册
	v1.POST("/signup", controllers.RegisterHandler)
	// 登录
	v1.POST("/login", controllers.LoginHandler)

	// 使用JWT中间件, 后续 /api/v1 操作均需要认证
	v1.Use(middlewares.JWTMiddleware())
	// 社区信息相关
	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
		v1.GET("/posts", controllers.GetPostListHandler)
	}

	// 测试登录状态
	r.GET("/ping", middlewares.JWTMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	// 404 page
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": http.StatusText(http.StatusNotFound)})
	})

	return r, nil
}
