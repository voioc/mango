package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/voioc/mango/lib/qiniu"
)

/*
 * @Author: Cedar
 * @Date: 2021-02-05 17:32:40
 * @LastEditors: Cedar
 * @LastEditTime: 2021-02-05 17:45:26
 * @FilePath: /Mango/app/handler/QiNiu.go
 */

// GetQiNiuToken 获取七牛云的上传token
func GetQiNiuToken(c *gin.Context) {

	token := qiniu.GetUploadToken()
	c.JSON(200, gin.H{"code": 0, "message": "success", "data": gin.H{"token": token}})
	return
}
