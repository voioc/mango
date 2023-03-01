package handler

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/voioc/mango/app/wx/define"
	"github.com/voioc/mango/common"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

//发送信息
type Content struct {
	Content string `json:"content"`
}

//开启回调模式验证
func WxAuth(c *gin.Context) {
	msgSignature := cast.ToString(c.Query("msg_signature"))
	//时间戳
	timestamp := cast.ToString(c.Query("timestamp"))
	//随机数
	nonce := cast.ToString(c.Query("nonce"))

	echostr, err := url.PathUnescape(c.Query("echostr"))
	if err != nil {
		fmt.Println("url解码失败")
		return
	}

	wxcpt := common.NewWXBizMsgCrypt(common.TOKEN, common.AESKEY, common.CORPID, common.XmlType)
	echoStr, cryptErr := wxcpt.VerifyURL(msgSignature, cast.ToString(timestamp), nonce, echostr)
	if nil != cryptErr {
		fmt.Println("verifyUrl fail", cryptErr)
		return
	}

	// id := x[appIDstart : int(appIDstart)+len(common.CORPID)]
	// if string(id) == common.CORPID {
	// 	return c.JSONBlob(200, x[20:20+length])
	// }
	c.String(http.StatusOK, string(echoStr))
	// return errors.New("微信验证appID错误, 微信请求值: " + string(id) + ", 配置文件内配置为: " + corpId)
}

func MsgBack(c *gin.Context) {
	/*
	   指令回调URL： 微信服务器推送suite_ticket以及安装应用时推送auth_code时。
	*/
	//企业微信加密签名
	msgSignature := cast.ToString(c.Query("msg_signature"))
	//时间戳
	timestamp := cast.ToString(c.Query("timestamp"))
	//随机数
	nonce := cast.ToString(c.Query("nonce"))
	// post请求的密文数据
	// defer c.Request.Body.Close()
	// con, _ := ioutil.ReadAll(c.Request.Body) //获取post的数据

	// // 访问应用和企业回调传不同的ID
	// wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(common.TOKEN, common.AESKEY, model.SuitId, wxbizmsgcrypt.XmlType)
	// msg, cryptErr := wxcpt.DecryptMsg(msgSignature, timestamp, nonce, con)
	// if nil != cryptErr {
	// 	fmt.Println("DecryptMsg fail", cryptErr)
	// 	return
	// }
	// fmt.Println(string(msg))

	// var content MsgContent
	// xml.Unmarshal(msg, &content)
	// var changeContent ChangeContent
	// xml.Unmarshal(msg, &changeContent)
	// fmt.Println(changeContent)

	defer c.Request.Body.Close()
	con, _ := ioutil.ReadAll(c.Request.Body) //获取post的数据
	fmt.Printf("con: %+v", string(con))

	// var textMsg define.WXTextMsg
	// err := c.ShouldBindXML(&textMsg)
	// if err != nil {
	// 	log.Printf("[消息接收] - XML数据包解析失败: %v\n", err)
	// 	return
	// }

	wxcpt := common.NewWXBizMsgCrypt(common.TOKEN, common.AESKEY, common.CORPID, common.XmlType)
	msg, err := wxcpt.DecryptMsg(msgSignature, timestamp, nonce, con)
	if err != nil {
		fmt.Println(err.ErrCode, err.ErrMsg)
	}

	fmt.Println(string(msg))

	var content define.MsgContent
	if err := xml.Unmarshal(msg, &content); err != nil {
		fmt.Println("反序列化错误")
	}

	fmt.Printf("%+v\n", content)

	//业务逻辑，根据信息需要进行的业务逻辑

	c.String(http.StatusOK, "success") //需要返回"success"不然企业微信认为此次请求错误
}
