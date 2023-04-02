package router

import (
	"github.com/voioc/mango/app/wx/handler"

	"github.com/gin-gonic/gin"
)

// InitRoute 设置路由
func DemoRouter(r *gin.RouterGroup) {
	// r.Use()

	// Simple group: v1
	// r.GET("/", api.Index)

	// 内部url
	wxGroup := r.Use() // 单独绑定中间件
	wxGroup.GET("/auth", handler.WxAuth)
	wxGroup.GET("/public/auth", handler.PublicMsg)
	wxGroup.POST("/auth", handler.MsgBack)
}
