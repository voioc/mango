package middlewares

import (
	"github.com/gin-gonic/gin"
	// "legitlab.letv.cn/letv_mobile_server/lej-go/lib/logcus"
)

/*
 * @Author: Cedar
 * @Date: 2021-01-07 10:30:41
 * @LastEditors: Cedar
 * @LastEditTime: 2021-01-28 11:58:54
 * @FilePath: /Mango/middlewares/cors.go
 */

// Cors 开启跨域函数
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		defer func() {
			if err := recover(); err != nil {
				// 	logcus.Print("Panic info is: %v", err)
				// 	logcus.Print("Panic info is: %s", debug.Stack())
			}
		}()

		c.Next()
	}
}
