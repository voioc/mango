package service

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/voioc/coco/proxy"
	"github.com/voioc/mango/app/wx/define"
	"github.com/voioc/mango/common"

	"github.com/gin-gonic/gin"
)

type ChatService struct {
	common.Base
}

// ChatS instance
func ChatS(c *gin.Context) *ChatService {
	return &ChatService{common.NewBase(c)}
}

func (s *ChatService) Send(msg string) (*define.Chat, error) {
	header := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer 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",
		// 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
	}

	msgMap := map[string]string{"role": "user", "content": msg}
	var message []map[string]string
	message = append(message, msgMap)

	data := map[string]interface{}{
		"model":       "gpt-3.5-turbo",
		"messages":    message,
		"max_tokens":  1024,
		"temperature": 0.8,
	}
	params, _ := jsoniter.Marshal(data)
	url := "https://api.openai.com/v1/chat/completions"

	tmp, err := proxy.SimpleClient(url, "POST", header, params)
	// s.SetDebug(1, "{Get data from cz88, url: %s?%s (%s)}", url, public.UrlEncode(params), public.TimeCost(startTime))
	if err != nil {
		return nil, fmt.Errorf("proxy info: %s", err.Error())
	}

	if tmp.StatusCode != 200 {
		return nil, fmt.Errorf("http code: %d | data: %s", tmp.StatusCode, string(tmp.Body))
	}

	result := define.Chat{}
	if err := jsoniter.Unmarshal(tmp.Body, &result); err != nil {
		return nil, fmt.Errorf("decode: %s | value: %s", err.Error(), string(tmp.Body))
	}

	return &result, nil
}
