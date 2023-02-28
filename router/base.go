package router

import (
	"github.com/gin-gonic/gin"
)

// InitRoute 设置路由
func InitRouter(engine *gin.Engine) {
	// SLB 用来探活服务的，不要删除，不要删除，不要删除！！！
	// ------------------------------------------------
	engine.GET("/v1/ping", func(c *gin.Context) {
		c.String(200, "Hello, China, Hello World")
	})

	// // 组件探活
	// engine.GET("/v1/active", func(c *gin.Context) {
	// 	data := map[string]interface{}{
	// 		"db":    db.Active(),
	// 		"cache": cache.Active(),
	// 	}
	// 	c.JSON(http.StatusOK, common.Success(c, data))
	// })
	// // -------------------------------------------------

	rootGroup := engine.Group("/")

	// demo 项目路由
	demoRouter := rootGroup.Group("/wx")
	DemoRouter(demoRouter)
}
