package main

import (
	"bluebell/config"
	"bluebell/dataset/mysql"
	"bluebell/dataset/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/routes"
	"context"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Go Web
func main() {
	redBgBlack := color.New(color.FgBlack, color.BgRed).PrintlnFunc()
	// 1. 加载配置
	if err := config.Init("./config", "config", "yaml"); err != nil {
		redBgBlack("init config failed, err:%v\n", err)
		panic(err)
	}

	// 2. 初始化日志
	if err := logger.Init(); err != nil {
		redBgBlack("init logger failed, err:%v\n", err)
		panic(err)
	}
	defer logger.Sync()

	// 3. 初始化mysql数据库
	if err := mysql.InitMysql(); err != nil {
		redBgBlack("init mysql failed, err:%v\n", err)
		panic(err)
	}
	defer mysql.CloseMysql()

	// 4. 初始化redis数据库
	if err := redis.InitRedis(); err != nil {
		redBgBlack("init redis failed, err:%v\n", err)
		panic(err)
	}
	defer redis.CloseRedis()

	// 初始化ID生成器
	if err := snowflake.Init(config.Cfg.StartTime, config.Cfg.MachineID); err != nil {
		redBgBlack("init snowflake failed, err:%v\n", err)
	}
	// 5. 注册路由
	route, err := routes.SetupRoutes()

	if err != nil {
		redBgBlack("setupRoutes failed, err:%v\n", err)
		panic(err)
	}
	// 6. 关机or重启
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Cfg.WebConfig.Port),
		Handler: route,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Fatal("listen", zap.Error(err))
		}
	}()
	// 显示网址
	color.Green("http://localhost:%d\n", config.Cfg.WebConfig.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown Error", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
