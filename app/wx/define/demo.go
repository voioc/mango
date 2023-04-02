package define

type DemoRequest struct {
	ID int
}

type CDATA struct {
	Value string `xml:",cdata"`
}

type XML struct {
	ToUsername   CDATA `xml:"ToUsername"`
	FromUsername CDATA `xml:"FromUsername"`
	CreateTime   int64 `xml:"CreateTime"`
	MsgType      CDATA `xml:"MsgType"`
	Content      CDATA `xml:"Content"`
}

//ToUserName	成员UserID
//FromUserName	企业微信CorpID
//CreateTime	消息创建时间（整型）
//MsgType	消息类型，此时固定为：text
//Content	文本消息内容,最长不超过2048个字节，超过将截断
type ReplyTextMsg struct {
	ToUsername   string `xml:"ToUserName"`
	FromUsername string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
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

type Chat struct {
	ID      string         `json:"id"`
	Object  string         `json:"object"`
	Created int            `json:"created"`
	Model   string         `json:"model"`
	Choices []ChatChoice   `json:"choices"`
	Usage   map[string]int `json:"usage"`
}

type ChatChoice struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}
