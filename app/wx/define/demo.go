package define

type DemoRequest struct {
	ID int
}

//ToUserName	成员UserID
//FromUserName	企业微信CorpID
//CreateTime	消息创建时间（整型）
//MsgType	消息类型，此时固定为：text
//Content	文本消息内容,最长不超过2048个字节，超过将截断
type ReplyTextMsg struct {
	ToUsername   string `xml:"ToUserName"`
	FromUsername string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
}
type MsgContent struct {
	ToUsername   string `xml:"ToUserName"`
	FromUsername string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	Msgid        string `xml:"MsgId"`
	Agentid      int    `xml:"AgentID"`
}

type WxVerify struct {
	ToUserName string
	Encrypt    string
	AgentID    int
}
