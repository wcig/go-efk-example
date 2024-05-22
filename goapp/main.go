package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	serviceName     = "goapp"
	logFieldService = "service_name"
	logFile         = "log/app.log"
)

var (
	logger *zap.SugaredLogger
)

func main() {
	initLog()
	runServer()
}

func initLog() {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "ts"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	w := io.MultiWriter(file, os.Stdout)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(w),
		zap.DebugLevel,
	)
	fields := zap.Fields(zap.String(logFieldService, serviceName))
	options := []zap.Option{fields, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)}
	logger = zap.New(core, options...).Sugar()
}

func runServer() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(accessLogger)
	router.GET("/ready", readyHandler)
	if err := router.Run(":38081"); err != nil {
		panic(err)
	}
}

func accessLogger(c *gin.Context) {
	start := time.Now()
	path := c.Request.Host + c.Request.URL.Path
	raw := c.Request.URL.RawQuery

	c.Next()

	end := time.Now()
	latency := end.Sub(start)
	clientIP := c.ClientIP()
	method := c.Request.Method
	statusCode := c.Writer.Status()
	if raw != "" {
		path += "?" + raw
	}
	logger.Infof("[GIN] %v | %3d | %13v | %15s | %-7s | %s",
		end.Format("2006/01/02 - 15:04:05"),
		statusCode,
		latency,
		clientIP,
		method,
		path,
	)
}

func readyHandler(c *gin.Context) {
	logger.Debug(">> ready debug log")
	logger.Info(">> ready info log")
	logger.Warn(">> ready warn log")
	logger.Error(">> ready error log")
	c.JSON(http.StatusOK, gin.H{"code": 0})
}
