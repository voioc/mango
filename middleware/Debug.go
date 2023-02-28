package middleware

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/voioc/coco/logzap"
	"github.com/voioc/coco/public"
	"github.com/voioc/mango/tool"
)

// AuthTK 接口请求验证
func AuthTK() gin.HandlerFunc {
	return func(c *gin.Context) {
		// base := strings.Split(c.Request.Host, ":") // 127.0.0.1:8001
		// host := base[0]
		// path := c.Request.URL.Path
		// fmt.Println(path)

		if !c.GetBool("_debug") { // 非本机请求，非debug请求
			tk := c.GetHeader("TK")
			text, secret := createSecret(c)
			version := c.Query("version")
			if strings.Contains(text, "com.sevenVideo.app.android") { // 不是本机请求
				if tk != secret {
					// ip := c.ClientIP()
					// public.FilePutContents("/tmp/token_go.log", fmt.Sprintf("%s | %s - %s - %s - %s", tk, secret, text, ip, version), true)
					public.FilePutContents("/tmp/token_go.log", fmt.Sprintf("ip: %s | tk: %s | secret: %s | str: %s | version: %s", c.ClientIP(), tk, secret, text, version), true)
					if version != "" && tool.VersionOrdinal(version) < tool.VersionOrdinal("1.9.8") {
						logzap.Ex(c, "mango", "illegal request %s", c.Request.URL.Path+"?"+c.Request.URL.Query().Encode())
						c.JSON(http.StatusOK, gin.H{"code": 2001, "msg": "非法请求", "data": "", "ext": ""})
						c.Abort()
						return
					}
				}
			}
		}

		c.Next()
	}
}

func createSecret(c *gin.Context) (string, string) {
	// tk := c.GetHeader("TK-Token")
	timeStamp := c.GetHeader("t")

	// path := c.Request.URL.Path
	params := c.Request.URL.Query()
	// fmt.Println(params)

	// 参数排序
	// To store the keys in slice in sorted order
	var keys []string
	for k := range c.Request.URL.Query() {
		if k != "_flush" && k != "_debug" {
			keys = append(keys, k)
		}

	}
	sort.Strings(keys)

	// To perform the operation you wantdddd
	// version := ""
	values := "" // 重置url顺序
	for _, k := range keys {
		// for _, row := range params[k] {
		// 	url += fmt.Sprintf("%s=%s&", k, row)	// 最后一个字符为&
		// }

		end := len(params[k]) - 1
		newVal := params[k][end]
		if k == "version" {
			// version = params[k][end]
		} else if k == "sj" {
			timeStamp = params[k][end]
		} else {
			newVal = url.QueryEscape(params[k][end])
		}

		values += newVal
	}

	text := values + timeStamp + "XSpeUFjJ"

	hash := md5.Sum([]byte(text))
	secret := fmt.Sprintf("%x", hash)

	return text, secret
}
