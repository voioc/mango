package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/voioc/coco/logzap"
	"github.com/voioc/coco/proxy"
	"github.com/voioc/mango/common"
)

//文本消息
type RobotText struct {
	MsgType string `json:"msgtype"`
	Text    `json:"text"`
	// At      `json:"at,omitempty"`
}

//文本消息的文本
type Text struct {
	Content string `json:"content"`
}

//机器人配置
const (
	webHook = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=1e6c5809-54ed-4566-abaf-53835b7748ff"
)

// PanicInterceptor Panic拦截器
func PanicInterceptor() gin.RecoveryFunc {
	return func(c *gin.Context, err interface{}) {
		trace := PanicTrace(2)
		// env := os.Getenv("envType")

		Message := RobotText{
			MsgType: "text",
			Text:    Text{Content: fmt.Sprintf("Panic异常!!: %v \n %s", err, string(trace))},
		}

		msg, _ := jsoniter.Marshal(Message)
		header := map[string]string{"Content-Type": "application/json"}

		//报警通知
		if _, err := proxy.SimpleClient(webHook, "POST", header, msg); err != nil {
			fmt.Println(err.Error())
			logzap.Ex(c.Request.Context(), "Panic", "proxy: %s", err.Error())
		}

		logzap.Ex(c.Request.Context(), "Panic", "PanicStack: %s", string(trace))
		c.JSON(http.StatusOK, common.Error(c, common.ERROR_INTER))
		c.Abort()
		return
	}
}

// 优化堆栈信息
func PanicTrace(kb int) []byte {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<10) //KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")
	return stack
}
