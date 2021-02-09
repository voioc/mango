/*
 * @Author: Cedar
 * @Date: 2021-02-09 09:59:11
 * @LastEditors: Cedar
 * @LastEditTime: 2021-02-09 16:36:17
 * @FilePath: /Mango/app/model/MVod.go
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

/*
 * @Author: Cedar
 * @Date: 2021-02-09 09:59:11
 * @LastEditors: Cedar
 * @LastEditTime: 2021-02-09 14:53:30
 * @FilePath: /Mango/app/model/MVod.go
 */
type Vods struct {
	Id          int64     `json:"id"`
	Category    string    `xorm:"category" json:"category"`
	SubCategory string    `xorm:"sub_category" json:"sub_category"`
	Title       string    `xorm:"title" json:"title"`
	SubTitle    string    `xorm:"sub_title" json:"sub_title"`
	Status      int       `xorm:"status" json:"status"`
	Reason      string    `xorm:"reason" json:"reason"`
	Color       string    `xorm:"color" json:"-"`
	Tag         string    `xorm:"tag" json:"-"`
	Thumb       string    `xorm:"thumb" json:"thumb"`
	Pics        string    `xorm:"pics" json:"pics"`
	Author      int       `xorm:"author" json:"author"`
	Remarks     string    `xorm:"remarks" json:"remarks"`
	Desc        string    `xorm:"desc" json:"desc"`
	Total       string    `xorm:"total" json:"total"`
	Area        string    `xorm:"area" json:"area"`
	Lang        string    `xorm:"lang" json:"lang"`
	Isend       string    `xorm:"isend" json:"isend"`
	Lock        string    `xorm:"lock" json:"lock"`
	Level       string    `xorm:"level" json:"level"`
	Hits        int       `xorm:"hits" json:"hits"`
	Duration    string    `xorm:"duration" json:"duration"`
	Score       float64   `xorm:"score" json:"score"`
	Publish     int       `xorm:"publish" json:"publish"`
	Sort        int       `xorm:"sort" json:"sort"`
	PlayUrl     string    `xorm:"play_url" json:"play_url"`
	DownUrl     string    `xorm:"down_url" json:"down_url"`
	CreatedId   int       `xorm:"created_id" json:"-"`
	CreatedAt   time.Time `xorm:"created_at" json:"-"`
	UpdatedId   int       `xorm:"updated_id" json:"-"`
	UpdatedAt   time.Time `xorm:"updated_at" json:"-"`
	DeleteId    int       `xorm:"deleted_id" json:"-"`
	DeletedAt   time.Time `xorm:"deleted_at" json:"-"`
}

// GetVodList 获取视频列表
func (bm *BaseModel) GetVodList(page string) *[]Vods {
	key := "vod_list_page_" + page
	var data []Vods
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
	if err := GetDB().Where("status = ?", 1).Limit(20, 20*pageInt).Find(&data); err != nil {
		logcus.Print("error", "Get data from mysql error: "+err.Error())
	} else {
		bm.SetDebug(fmt.Sprintf("{Get data from Mysql ("+public.TimeCost(StartTime)+"): %s}", data), 1)
		cache.SetCache(key, data, -1)
	}

	return &data
}
