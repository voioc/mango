package common

/*
 * @Author: Cedar
 * @Date: 2020-11-06 14:22:30
 * @LastEditors: Cedar
 * @LastEditTime: 2020-11-06 16:47:57
 * @FilePath: /LeView/app/common/Base.go
 */

import (
	"runtime"
	"strconv"
	"strings"
)

// Base 基类
type Base struct {
	Debug   *[]string
	IsCache bool
}

// NewBase NewBase
func NewBase() *Base {
	return &Base{}
}

// SetIsCache 是否使用缓存
func (base *Base) SetIsCache(isCache bool) {
	base.IsCache = isCache
}

// SetDebugP 设置debug的指针
func (base *Base) SetDebugP(debug *[]string) {
	base.Debug = debug
}

// SetDebug 写入debug信息
func (base *Base) SetDebug(str string, depth int) {
	if base.Debug != nil {
		if depth == 0 {
			depth = 1
		}

		_, file, line, _ := runtime.Caller(depth)
		path := strings.LastIndexByte(file, '/')
		tmp := string([]byte(file)[path+1:]) + "(line " + strconv.Itoa(line) + "): " + str
		*base.Debug = append(*base.Debug, tmp)
	}
}
