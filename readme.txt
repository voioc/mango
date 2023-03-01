package bot

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"wxbot-test/wxbizmsgcrypt"
)

//ToUserName	成员UserID
//FromUserName	企业微信CorpID
//CreateTime	消息创建时间（整型）
//MsgType	消息类型，此时固定为：text
//Content	文本消息内容,最长不超过2048个字节，超过将截断
type ReplyTextMsg struct {
	ToUsername   string `xml:"ToUserName"`
	FromUsername string `xml:"FromUserName"`
	CreateTime   uint32 `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
}
type MsgContent struct {
	ToUsername   string `xml:"ToUserName"`
	FromUsername string `xml:"FromUserName"`
	CreateTime   uint32 `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	Msgid        string `xml:"MsgId"`
	Agentid      uint32 `xml:"AgentId"`
}

var (
	corpId, token, encodingAesKey string
	wxcrypt                       *wxbizmsgcrypt.WXBizMsgCrypt
)

func init() {
	// 读取配置文件
	corpId = "ww592d57xxxxxxx"
	token = "VbmqbkBQWnxxxxxxxxx"
	encodingAesKey = "HDyFtR7DB6oyeG7DuzCOgjWg7xxxxxxxxxxxxxxxxx"
	// receive_id 企业应用的回调，表示corpid
	wxcrypt = wxbizmsgcrypt.NewWXBizMsgCrypt(token, encodingAesKey, corpId, wxbizmsgcrypt.XmlType)
}

func Start() {
	// 开启一个http服务器，接收来自企业微信的消息
	http.HandleFunc("/api/bot/message", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Println("接收到验证请求")
			handleVerify(w, r)
		} else if r.Method == "POST" {
			fmt.Println("接收到消息")
			handleMessage(w, r)
		}
	})
	log.Fatalln(http.ListenAndServe("127.0.0.1:8888", nil))
}
func handleMessage(w http.ResponseWriter, r *http.Request) {
	msgSignature := r.URL.Query().Get("msg_signature")
	timestamp := r.URL.Query().Get("timestamp")
	nonce := r.URL.Query().Get("nonce")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("读取Body错误", err.Error())
	} else {
		msg, err := wxcrypt.DecryptMsg(msgSignature, timestamp, nonce, body)
		if err != nil {
			fmt.Println("解密消息Body错误", err.ErrMsg)
		} else {
			var msgContent MsgContent
			err := xml.Unmarshal(msg, &msgContent)
			if err != nil {
				fmt.Println("反序列化错误")
			} else {
				fmt.Println(msgContent)
				// 回复信息
				replyMsg, _ := xml.Marshal(ReplyTextMsg{
					ToUsername:   msgContent.FromUsername,
					FromUsername: msgContent.ToUsername,
					CreateTime:   msgContent.CreateTime,
					MsgType:      "text",
					Content:      "Receive",
				})
				encryptMsg, cryptErr := wxcrypt.EncryptMsg(string(replyMsg), timestamp, nonce)
				if cryptErr != nil {
					fmt.Println("回复加密出错", cryptErr)
				} else {
					fmt.Println(string(encryptMsg))
					l, err := w.Write(encryptMsg)
					if err != nil {
						fmt.Println("返回消息失败")
					} else {
						fmt.Println("成功写入", l)
					}
				}

			}
		}
	}
}
func handleVerify(w http.ResponseWriter, r *http.Request) {
	msgSignature := r.URL.Query().Get("msg_signature")
	timestamp := r.URL.Query().Get("timestamp")
	nonce := r.URL.Query().Get("nonce")
	echoStr := r.URL.Query().Get("echostr")
	// 合法性验证
	echoStrBytes, err := wxcrypt.VerifyURL(msgSignature, timestamp, nonce, echoStr)
	if err != nil {
		fmt.Println("验证失败", err.ErrMsg)
	} else {
		fmt.Println("验证成功", string(echoStrBytes))
		// 需要返回才能通过验证
		_, err := w.Write(echoStrBytes)
		if err != nil {
			fmt.Println("返回验证结果失败", err.Error())
		}
	}
}