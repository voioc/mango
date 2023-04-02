package handler

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/voioc/coco/proxy"
	"github.com/voioc/mango/app/wx/define"
)

//开启回调模式验证
func PublicMsg(c *gin.Context) {
	// fmt.Println(WxGetAccessToken())
	// echoStr, err := url.PathUnescape(c.Query("echostr"))
	// if err != nil {
	// 	fmt.Println("url解码失败")
	// 	return
	// }

	// c.String(http.StatusOK, string(echoStr))

	defer c.Request.Body.Close()
	con, _ := ioutil.ReadAll(c.Request.Body) //获取post的数据
	fmt.Printf("con: %+v", string(con))

	var content define.MsgContent
	if err := xml.Unmarshal(con, &content); err != nil {
		fmt.Println("反序列化错误")
		return
	}

	fmt.Printf("%+v\n", content)

	// chat, err2 := service.ChatS(c).Send(content.Content)
	// if err2 != nil {
	// 	fmt.Println(err2.Error())
	// 	return
	// }

	// fmt.Printf("%+v\n", chat)
	replyContent := "I don't know"
	// if len(chat.Choices) > 0 {
	// replyContent = chat.Choices[0].Text
	// }

	// 回复信息
	reply, _ := xml.Marshal(define.ReplyText{
		XMLName:      xml.Name{Local: "xml"},
		ToUsername:   define.CDATA{Value: content.FromUsername},
		FromUsername: define.CDATA{Value: content.ToUsername},
		CreateTime:   time.Now().Unix(),
		MsgType:      define.CDATA{Value: "text"},
		Content:      define.CDATA{Value: replyContent},
	})

	// encryptMsg, cryptErr := wxcpt.EncryptMsg(string(reply), timestamp, nonce)
	// if cryptErr != nil {
	// 	fmt.Println("回复加密出错", cryptErr)
	// 	return
	// }
	fmt.Println("reply 1111111", string(reply))
	reply = []byte("<xml><ToUserName><![CDATA[oDk236LGfpuPCzAqH09I9RzFYw1c]]></ToUserName><FromUserName><![CDATA[gh_2680178c02e1]]></FromUserName><CreateTime>1680448101</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[" + replyContent + "]]></Content></xml>")
	// reply = []byte("<xml><ToUsername><![CDATA[oDk236LGfpuPCzAqH09I9RzFYw1c]]></ToUsername><FromUsername><![CDATA[gh_2680178c02e1]]></FromUsername><CreateTime>1680449116</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[你好]]></Content></xml>")
	fmt.Println("reply encry", string(reply))
	if num, err := c.Writer.Write(reply); err != nil {
		fmt.Println("返回消息失败: ", err.Error())
		return
	} else {
		fmt.Println("success: ", num)
	}

	// accessToken := WxGetAccessToken()
	// params, _ := jsoniter.Marshal(reply)
	// url := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=" + accessToken
	// tmp, _ := proxy.SimpleClient(url, "POST", nil, params)
	// fmt.Println(string(tmp.Body))

	//业务逻辑，根据信息需要进行的业务逻辑
	c.String(http.StatusOK, "success") //需要返回"success"不然企业微信认为此次请求错误
	// <xml><ToUserName><![CDATA[oDk236LGfpuPCzAqH09I9RzFYw1c]]></ToUserName><FromUserName><![CDATA[gh_2680178c02e1]]></FromUserName><CreateTime>1680448101</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[你好]]></Content></xml>
	// <xml><ToUsername><![CDATA[oDk236LGfpuPCzAqH09I9RzFYw1c]]></ToUsername><FromUsername><![CDATA[gh_2680178c02e1]]></FromUsername><CreateTime>1680449116</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[你好]]></Content></xml>
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
	if err := server.Send(); err != nil {
		fmt.Println(err.Error())
	}
}

// WxGetAccessToken 获取微信accesstoken
func WxGetAccessToken() string {
	params := map[string]string{
		"grant_type": "client_credential",
		"appid":      "wx2af03ead301cf223",
		"secret":     "b98b0b0719aa6042f4ac28d7d504d110",
	}
	url := "https://api.weixin.qq.com/cgi-bin/token"

	tmp, _ := proxy.SimpleClient(url, "GET", nil, params)
	token := jsoniter.Get(tmp.Body, "access_token").ToString()

	return token
}
