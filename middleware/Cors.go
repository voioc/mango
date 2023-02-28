/*
 * @Description: Do not edit
 * @Author: Jianxuesong
 * @Date: 2021-06-01 15:57:55
 * @LastEditors: Jianxuesong
 * @LastEditTime: 2021-06-16 15:14:24
 * @FilePath: /Melon/middleware/Cors.go
 */
package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("start", time.Now())
		c.Set("_debug", c.Query("_debug") == "1")
		// c.Set("_flush", true)
		c.Set("_flush", c.Query("_flush") == "1")

		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			// 可将将* 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// defer func() {
		// 	if err := recover(); err != nil {
		// 		logcus.Error(fmt.Sprintf("Cos Panic info is: %v", err))
		// 	}
		// }()

		c.Next()

		//处理后获取消耗时间
		// costTime := time.Since(c.GetTime("start"))
		// url := c.Request.URL.String()
		// // c.GetTime()
		// // fmt.Printf("the request URL %s cost %v\n", url, costTime)
	}
}
