package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

//开启回调模式验证
func PublicMsg(c *gin.Context) {
	// echoStr, err := url.PathUnescape(c.Query("echostr"))
	// if err != nil {
	// 	fmt.Println("url解码失败")
	// 	return
	// }

	// c.String(http.StatusOK, string(echoStr))

	defer c.Request.Body.Close()
	con, _ := ioutil.ReadAll(c.Request.Body) //获取post的数据
	fmt.Printf("con: %+v", string(con))
	handleMsg(c.Writer, c.Request)
}

func handleMsg(rw http.ResponseWriter, req *http.Request) {
	wc := wechat.NewWechat()
	// 这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:     "wx2af03ead301cf223",
		AppSecret: "",
		Token:     "bHzz89fXlVBJZllDpVgeHw7ymmMLoHU",
		//EncodingAESKey: "xxxx",
		Cache: memory,
	}
	officialAccount := wc.GetOfficialAccount(cfg)

	// 传入request和responseWriter
	server := officialAccount.GetServer(req, rw)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		//TODO
		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}
