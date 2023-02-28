package tool

import (
	"fmt"
	"strconv"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	jsoniter "github.com/json-iterator/go"
)

func GetCacheKey(key ...string) string {
	prefix := []string{"qiqi"}
	key = append(prefix, key...)
	return strings.Join(key, ":")
}

func RuneToString(data string) string {
	result := ""

	tmp := strings.Split(data, "|")
	for _, data := range tmp {
		son_result := ""
		if data != "" {
			for _, row := range data {
				if son_result == "" {
					son_result = strconv.Itoa(int(row))
				} else {
					son_result += "_" + strconv.Itoa(int(row))
				}
			}
		}

		if result == "" {
			result = son_result
		} else {
			result += "#" + son_result
		}
	}

	return result
}

func SendSMS() error {
	accessKeyID := "LTAI4FzEE8zMMXxetMLciBEY"
	accessKeySecret := "dHvprsZ48rzAK3R7zspry1grXFLve1"
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &accessKeyID,
		// 您的AccessKey Secret
		AccessKeySecret: &accessKeySecret,
	}

	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")

	// client = &dysmsapi20170525.Client{}
	client, err := dysmsapi20170525.NewClient(config)
	if err != nil {
		return err
	}

	phone := "15210232278"
	SignName := "御风云"
	TemplateCode := "SMS_210180005"
	param, _ := jsoniter.Marshal(map[string]int{"code": 1234})
	templateParam := string(param)
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{PhoneNumbers: &phone, SignName: &SignName, TemplateCode: &TemplateCode, TemplateParam: &templateParam}
	// 复制代码运行请自行打印 API 的返回值
	if result, err := client.SendSms(sendSmsRequest); err != nil {
		return err
	} else {
		fmt.Println(result)
	}

	return err
}

func VersionOrdinal(version string) string {
	// ISO/IEC 14651:2011
	const maxByte = 1<<8 - 1
	vo := make([]byte, 0, len(version)+8)
	j := -1
	for i := 0; i < len(version); i++ {
		b := version[i]
		if '0' > b || b > '9' {
			vo = append(vo, b)
			j = -1
			continue
		}
		if j == -1 {
			vo = append(vo, 0x00)
			j = len(vo) - 1
		}
		if vo[j] == 1 && vo[j+1] == '0' {
			vo[j+1] = b
			continue
		}
		if vo[j]+1 > maxByte {
			panic("VersionOrdinal: invalid version")
		}
		vo = append(vo, b)
		vo[j]++
	}
	return string(vo)
}
