/*
 * @Author: Cedar
 * @Date: 2021-02-09 09:59:11
 * @LastEditors: Cedar
 * @LastEditTime: 2021-02-09 17:48:58
 * @FilePath: /Mango/app/model/MVodUser.go
 */
package model

import (
	"fmt"
	"strconv"
	"time"

	"github.com/voioc/coco/cache"
	"github.com/voioc/coco/logcus"
	"github.com/voioc/coco/public"
)

type VodUser struct {
	Vods   `xorm:"extends"`
	Author Users `xorm:"extends" json:"author"`
}

func (VodUser) TableName() string {
	return "vods"
}

// GetVodAuthorList 获取视频列表
func (bm *BaseModel) GetVodAuthorList(page string) *[]VodUser {
	key := "vod_user_list_page_" + page
	var data []VodUser
	StartTime := time.Now()

	if !bm.GetBool("_flush") {
		if ok, err := cache.GetCache(key, &data); ok && err == nil { // 远程缓存中取出数据并正确解析则返回
			if len(data) > 0 {
				bm.SetDebug(fmt.Sprintf("{Get data from local cache Key: %s ("+public.TimeCost(StartTime)+")}", key), 1)
				return &data
			}
		}
	}

	StartTime = time.Now()
	pageInt, _ := strconv.Atoi(page)
	// if err := GetDB().Where("status = ?", 1).Limit(20, 20*pageInt).Find(&data); err != nil {
	if err := GetDB().Where("vods.status = ?", 1).Limit(20, (pageInt-1)*20).Join("LEFT OUTER", "users", "users.id = vods.author").Find(&data); err != nil {
		logcus.Print("error", "Get data from mysql error: "+err.Error())
	} else {
		bm.SetDebug(fmt.Sprintf("{Get data from Mysql ("+public.TimeCost(StartTime)+"): %s}", data), 1)
		cache.SetCache(key, data, -1)
	}

	return &data
}
