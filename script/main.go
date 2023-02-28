package main

import (
	"log"

	"github.com/voioc/mango/script/demo"

	"github.com/robfig/cron/v3"
)

func main() {
	log.Println("Starting...")

	//新建一个定时任务对象
	c := cron.New()

	c.AddFunc("@every 0h0m1s", demo.DemoS.China())
	c.AddJob("@every 0h0m1s", &demo.DemoS)

	// 启动计划任务
	c.Start()

	// 关闭计划任务, 但是不能关闭已经在执行中的任务.
	// defer c.Stop()

	select {}
}

// func main() {

// 	channel := flag.String("t", "0", "输入类型")
// 	flag.Parse()

// 	switch *channel {
// 	case "1":
// 		fmt.Println("now: ", time.Now().Format("2006:01:02 15:04:05"))
// 	case "2":
// 		fmt.Println("Hello World")
// 	case "3":
// 		fmt.Println("Hello China")
// 	default:
// 		fmt.Println("参数为空")
// 	}
// }
