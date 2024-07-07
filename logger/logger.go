package logger

import (
	"bluebell/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"reflect"
	"runtime/debug"
	"strings"
	"time"
)

// Init 初始化Logger
func Init() (err error) {
	writeSyncer := getLogWriter(
		config.Cfg.Filename,
		config.Cfg.MaxSize,
		config.Cfg.MaxBackups,
		config.Cfg.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(config.Cfg.Level))
	if err != nil {
		return err
	}
	var core zapcore.Core
	if strings.ToLower(config.Cfg.Mode) == "dev" {
		// 开发模式, 日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(WrappedWriteSyncer{os.Stdout}), zapcore.DebugLevel))
	} else {
		// 发布模式, 日志仅输出到文件
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	webLogger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(webLogger) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	return err
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: err check
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

// WrappedWriteSyncer 取消了Sync方法实际功能
type WrappedWriteSyncer struct {
	file *os.File
}

func (mws WrappedWriteSyncer) Write(p []byte) (n int, err error) {
	return mws.file.Write(p)
}

func (mws WrappedWriteSyncer) Sync() (err error) {
	return err
}

func Sync() {
	err := zap.L().Sync()
	if err != nil {
		t := reflect.TypeOf(err)
		fmt.Printf("zap sync err: [%s] %s\n", t.String(), err.Error())
	}
}
