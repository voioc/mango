package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	// "legitlab.letv.cn/letv_mobile_server/lej-go/lib/logcus"
)

/*
 * @Author: Cedar
 * @Date: 2021-01-07 10:30:41
 * @LastEditors: Cedar
 * @LastEditTime: 2021-02-09 15:41:40
 * @FilePath: /Mango/middlewares/params.go
 */

// Params 开启跨域函数
func Params(c *gin.Context) {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		// 	logcus.Print("Panic info is: %v", err)
	// 		// 	logcus.Print("Panic info is: %s", debug.Stack())
	// 	}
	// }()
	if _debug := c.DefaultQuery("_debug", "0"); _debug == "1" {
		c.Set("_debug", &[]string{})
		c.Set("start", time.Now())
	}
	if _flush := c.DefaultQuery("_flush", "0"); _flush == "1" {
		c.Set("_flush", true)
	}

	c.Next()
}
