package main

/*
 * @Author: Cedar
 * @Date: 2020-11-05 11:03:52
 * @LastEditors: Cedar
 * @LastEditTime: 2021-02-09 16:35:03
 * @FilePath: /Mango/mango.go
 */

import (
	"io"
	"log"
	"os"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/voioc/coco/cache"
	"github.com/voioc/mango/app/build"

	"github.com/gin-gonic/gin"
	"github.com/voioc/coco/config"
	"github.com/voioc/mango/app/handler"
	"github.com/voioc/mango/middlewares"
)

func main() {
	config.LoadConfig(build.GetConfigFile())

	// // 初始化错误日志
	// ef := config.GetConfig().GetString("log.error") // error log file
	// el := config.GetConfig().GetString("log.level") // error log level
	// logcus.Init(ef, el)
	cache.Init()

	// 初始化访问日志
	alf := config.GetConfig().GetString("log.access")
	access, err := os.OpenFile(alf, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}

	defer access.Close()
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(config.GetConfig().GetString("log.access"))
	gin.DefaultWriter = io.MultiWriter(access)

	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.ForceConsoleColor()
	gin.DefaultWriter = io.MultiWriter(access, os.Stdout)

	r := gin.Default()

	r.Use(middlewares.Params)

	r.Use(middlewares.Cors())
	{
		r.GET("/login", handler.Login)
		r.GET("/post/index", handler.PostIndex)
		r.GET("/qiniu/token", handler.GetQiNiuToken)
		r.GET("/vod/index", handler.VodIndex)
	}

	// r.GET("/task/do", handler.TaskDo)
	// r.GET("/task/finish", handler.TaskFinish)
	// r.Run() // listen and serve on 0.0.0.0:8080

	log.Println(gin.Mode())

	server := endless.NewServer(":8002", r)
	server.BeforeBegin = func(add string) {
		log.Printf("pid is %d", syscall.Getpid())
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("server err: %v", err)
	}
}
