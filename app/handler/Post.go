package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/voioc/mango/app/model"
)

/*
 * @Author: Cedar
 * @Date: 2021-01-28 11:08:27
 * @LastEditors: Cedar
 * @LastEditTime: 2021-01-28 11:29:01
 * @FilePath: /Mango/app/handler/Post.go
 */

func PostIndex(c *gin.Context) {
	data := []model.Posts{}
	if err := model.GetDB().Where("status = ? ", 1).Find(&data); err != nil {
		fmt.Println(err.Error())
		// logcus.Print("error", "Get data from mysql error: "+err.Error())
	} else {
		// bm.SetDebug(fmt.Sprintf("{Get data from Mysql ("+public.TimeCost(StartTime)+"): %s}", data), 1)
		// cache.SetCache(key, data, -1)
	}

	c.JSON(200, gin.H{"code": 0, "message": "success", "data": data})
	return
}
