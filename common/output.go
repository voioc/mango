package common

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

type Page struct {
	CurPage int `json:"cur_page"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}

// Error 通用错误的错误返回
func Error(c *gin.Context, code int, message ...string) gin.H {
	msg := ""
	if mm, flag := INFO[code]; flag {
		msg = mm
	}

	if len(message) > 0 {
		msg = message[0]
	}

	return SetOutput(c, code, msg, nil, nil)
}

// Success 通用的成功返回
func Success(c *gin.Context, params ...interface{}) gin.H {
	var data, ext interface{}
	if len(params) > 0 {
		data = params[0]
	}

	if len(params) > 1 {
		ext = params[1]
	}

	return SetOutput(c, STATUS_OK, "success", data, ext)
}

// SetOutput 自定义返回结构，自己构造
func SetOutput(c *gin.Context, code int, msg string, data, ext interface{}, params ...map[string]interface{}) gin.H {
	// RealCode := http.StatusOK
	result := gin.H{"code": code, "msg": msg, "data": data}

	if data != nil {
		result["data"] = data
	}

	if ext != nil {
		result["ext"] = ext
	}

	// 写入traceid
	result["x_trace_id"] = c.GetString("x_trace_id")

	// 兼容多余参数
	if len(params) > 0 {
		for key, value := range params[0] {
			result[key] = value
		}
	}

	base := Base{c}
	base.SetDebug(1, "[ALL] cost: %s)", TimeCost(c.GetTime("start")))

	if c.GetBool("_debug") {
		result["_debug"] = c.GetStringSlice("debug")
	}

	output, _ := jsoniter.MarshalToString(result)
	c.Set("output", output)

	return result
}

func TimeCost(start time.Time) string {
	tc := time.Since(start)
	return fmt.Sprintf("%v", tc)
}
