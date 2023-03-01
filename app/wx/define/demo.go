package define

type DemoRequest struct {
	ID int
}

// WXTextMsg 微信文本消息结构体
type WXTextMsg struct {
	ToUserName string
	Encrypt    string
	AgentID    int
}
