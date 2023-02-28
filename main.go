/*
 * @Description: Do not edit
 * @Author: Jianxuesong
 * @Date: 2021-05-18 14:28:40
 * @LastEditors: Jianxuesong
 * @LastEditTime: 2021-06-30 18:21:23
 * @FilePath: /Melon/main.go
 */
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"github.com/voioc/coco/logzap"

	"github.com/voioc/mango/middleware"
	"github.com/voioc/mango/router"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

var (
	// Version should be updated by hand at each release
	RunEnv      string
	ProjectPath string
	AppVersion  string
	GitCommit   string
	BuildTime   string
	GoVersion   string
)

var env string = "dev"

func init() {
	viper.SetDefault("env", env)

	// 获取当前环境变量
	if realEnv := strings.ToLower(os.Getenv("env")); realEnv != "" {
		env = realEnv
		viper.SetDefault("env", env)
	}

	// 默认根据环境变量来确定配置文件
	path, _ := filepath.Abs(filepath.Dir(""))        // 获取当前路径
	conf := path + "/config/config_" + env + ".toml" // 拼接配置文件

	versionFlag := flag.Bool("V", false, "print the version")
	configFile := flag.String("c", conf, "配置文件路径") // 手动置顶配置文件
	flag.Parse()

	if *versionFlag {
		fmt.Printf("App Version: %s \n", AppVersion)
		fmt.Printf("Git Commit: %s \n", GitCommit)
		fmt.Printf("Build Time: %s \n", BuildTime)
		fmt.Printf("Go Version: %s \n", GoVersion)
		os.Exit(0)
	}

	viper.SetConfigFile(*configFile) // 读取配置文件
	fmt.Println("Loading config file " + *configFile)

	if err := viper.ReadInConfig(); err != nil { //是否读取成功
		log.Fatalln("打开配置文件失败", err)
	}

	// 根据配置文件指定环境变量
	viper.SetDefault("env", viper.GetString("app.env"))

	// db.InitDB()  // 初始化数据库连接
	// cache.Init() // 初始化缓存连接

	logzap.InitZap(viper.GetString("log.error")) // 初始化日志组件
}

func main() {
	// logzap.InitZap()

	// 创建日志文件并设置为 gin.DefaultWriter
	// logPath := viper.GetString("log.access")
	// // logPath2 := build.GetConfig().GetString("log.access")
	// // fmt.Println(logPath, logPath2)
	// logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalln("打开日志文件失败：", err)
	// }

	// if viper.GetString("env") == "release" {
	// 	gin.SetMode(gin.ReleaseMode)
	// 	gin.DefaultWriter = io.MultiWriter(logFile)
	// } else {
	// 	gin.SetMode(gin.DebugMode)
	// 	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	// }

	// 日志文件不需要颜色
	//  gin.DisableConsoleColor()

	r := gin.New()
	// pprof.Register(r)
	// r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	// 自定义日志格式
	// 	return fmt.Sprintf("[%s] - %s \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.TimeStamp.Format(time.RFC3339),
	// 		param.ClientIP,
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	// }))

	// if err := tool.SendSMS(); err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println("短信发送成功")
	// }

	// 中间件
	r.Use(gin.Recovery(), middleware.Trace(), middleware.CORS(), middleware.ZapLogger())
	r.Use(gin.CustomRecovery(middleware.PanicInterceptor()))

	// 主页
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Melon")
	})
	// r.GET("/test", handler.TextIndex)

	router.InitRouter(r)

	fmt.Println("The service is running...")
	endless.ListenAndServe(":8001", r)

	// srv := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: r,
	// }

	// go func() {
	// 	// service connections
	// 	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()

	// // Wait for interrupt signal to gracefully shutdown the server with
	// // a timeout of 5 seconds.
	// quit := make(chan os.Signal)
	// // kill (no param) default send syscanll.SIGTERM
	// // kill -2 is syscall.SIGINT
	// // kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// <-quit
	// log.Println("Shutdown Server ...")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Shutdown:", err)
	// }
	// // catching ctx.Done(). timeout of 5 seconds.
	// select {
	// case <-ctx.Done():
	// 	log.Println("timeout of 5 seconds.")
	// }
	// log.Println("Server exiting")
}
