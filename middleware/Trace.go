package middleware

import (
    "context"
    "time"

    "github.com/gin-gonic/gin"
    uuid "github.com/satori/go.uuid"
    "github.com/spf13/viper"
    "github.com/voioc/coco/logzap"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

// Trace 接口请求验证
func Trace() gin.HandlerFunc {
    return func(c *gin.Context) {

        traceID := c.GetHeader("x_trace_id")
        if traceID == "" {
            traceID = uuid.NewV4().String()
        }

        ctx := context.WithValue(c.Request.Context(), logzap.ContextKey("x_trace_id"), traceID)
        c.Request = c.Request.WithContext(ctx)
        c.Set("x_trace_id", traceID)

        c.Next()
    }
}

// ZapLogger 接收gin框架默认的日志
func ZapLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now().Format("2006-01-02 15:04:05")
        c.Next()

        // data :=
        // output, _ := jsoniter.MarshalToString(gin.H(c.GetStringMap("output")))

        logInfo := []zapcore.Field{
            zap.String("time", t),
            zap.String("host", c.Request.Host),
            zap.String("remote_ip", c.ClientIP()),
            zap.String("method", c.Request.Method),
            zap.String("request_ur", c.Request.URL.Path+"?"+c.Request.URL.RawQuery),
            zap.Int("status", c.Writer.Status()),
            zap.String("user-agent", c.Request.UserAgent()),
            zap.String("type", "access"),
            // zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
            zap.String("x_trace_id", c.GetString("x_trace_id")),
        }

        if viper.GetBool("log.return") {
            logInfo = append(logInfo, zap.String("output", c.GetString("output")))
        }

        logzap.Zap().Info("request info", logInfo...)
    }
}
