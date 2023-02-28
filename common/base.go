package common

import (
    "context"
    "fmt"
    "runtime"
    "strconv"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
)

type Base struct {
    context context.Context
}

func NewBase(c context.Context) Base {
    return Base{c}
}

// Flush Flush
func (bm *Base) Flush() bool {
    if c, flag := interface{}(bm).(*gin.Context); flag {
        return c.GetBool("_flush")
    }

    return false
}

// Debug debug
func (bm *Base) Debug() bool {
    if c, flag := interface{}(bm.context).(*gin.Context); flag {
        return c.GetBool("_debug")
    }

    return false
}

func (bm *Base) Context() context.Context {
    if c, flag := interface{}(bm.context).(*gin.Context); flag {
        return c.Request.Context()
    }

    return bm.context
}

// SetDebug 设置debug
func (bm *Base) SetDebug(depth int, message string, a ...interface{}) {
    var c *gin.Context
    if tmpC, flag := interface{}(bm.context).(*gin.Context); !flag {
        return
    } else {
        c = tmpC
    }

    if depth == 0 {
        depth = 1
    }

    _, file, line, _ := runtime.Caller(depth)
    path := strings.LastIndexByte(file, '/')

    info := fmt.Sprintf(message, a...)
    tmp := string([]byte(file)[path+1:]) + "(line " + strconv.Itoa(line) + "): " + info

    debug := c.GetStringSlice("debug")
    c.Set("debug", append(debug, tmp))
}

// TimeCost 计算花费时间
func (bm *Base) TimeCost(t interface{}) string {
    var c *gin.Context
    if tmpC, flag := interface{}(bm.context).(*gin.Context); !flag {
        return ""
    } else {
        c = tmpC
    }

    cost := ""
    var start time.Time
    if key, flag := t.(string); flag {
        if start = c.GetTime(key); start.IsZero() {
            c.Set(key, time.Now())
        }
    }

    if newStart, flag := t.(time.Time); flag {
        start = newStart
    }

    tc := time.Since(start)
    cost = fmt.Sprintf("%v", tc)

    return cost
}
