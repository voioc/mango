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

	data := map[string]interface{}{
		"model":       "text-davinci-003",
		"prompt":      msg,
		"max_tokens":  1024,
		"temperature": 0.8,
	}
	params, _ := jsoniter.MarshalToString(data)
	url := "https://api.openai.com/v1/completions"

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
