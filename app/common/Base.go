package common

/*
 * @Author: Cedar
 * @Date: 2020-11-06 14:22:30
 * @LastEditors: Cedar
 * @LastEditTime: 2021-02-09 16:29:10
 * @FilePath: /Mango/app/common/Base.go
 */

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/voioc/coco/public"
)

// Base 基类
type Base struct {
	*gin.Context
}

// NewBase NewBase
func NewBase(c *gin.Context) *Base {
	return &Base{c}
}

// SetDebug 写入debug信息
func (b *Base) SetDebug(str string, depth int) {
	if _, ok := b.Get("_debug"); ok {
		if depth == 0 {
			depth = 1
		}

		_, file, line, _ := runtime.Caller(depth)
		path := strings.LastIndexByte(file, '/')
		tmp := string([]byte(file)[path+1:]) + "(line " + strconv.Itoa(line) + "): " + str

		b.Set("_debug", append(b.GetStringSlice("_debug"), tmp))
		// *base.Debug = append(*base.Debug, tmp)
	}
}

// GetOutPut 获取返回数据
func (b *Base) GetOutPut(data interface{}) *gin.H {
	result := gin.H{"code": 0, "message": "success"}

	if data == nil {
		data = map[string]string{}
	}

	result["data"] = data

	if _, ok := b.Get("_debug"); ok {
		debug := b.GetStringSlice("_debug")
		result["debug"] = append(debug, "[ALL] cost:"+public.TimeCost(b.GetTime("start")))
	}

	return &result
}
