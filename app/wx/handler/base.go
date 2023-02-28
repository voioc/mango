package handler

import (
	"fmt"
	"net/http"
	"net/url"

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

	wxcpt := common.NewWXBizMsgCrypt(common.TOKEN, common.AESKEY, common.CORPID, common.JsonType)
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
