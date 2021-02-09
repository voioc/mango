package model

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/voioc/coco/cache"
	"github.com/voioc/coco/logcus"
	"github.com/voioc/coco/public"
	"github.com/voioc/mango/app/common"
)

/*
 * @Author: Cedar
 * @Date: 2020-11-13 14:33:01
 * @LastEditors: Cedar
 * @LastEditTime: 2020-11-19 20:52:40
 * @FilePath: /Lej/models/BaseModel.go
 */

// BaseModel 基类
type BaseModel struct {
	common.Base
}

// NewBaseModel 基础
func NewBaseModel(c *gin.Context) *BaseModel {
	bm := BaseModel{common.Base{c}}

	// if debug != nil {
	// 	bm.SetDebugP(debug)
	// }

	// if nocache {
	// 	bm.SetNoCache(nocache)
	// }

	return &bm
}

// // CreateOrUpdate 插入或者更新
// func (bm *BaseModel) CreateOrUpdate(params map[string]interface{}, data interface{}, keys []string) error {
// 	where := ""
// 	value := make([]interface{}, 0, len(params))
// 	for col, val := range params {
// 		if where == "" {
// 			where += col + " = ?"
// 		} else {
// 			where += "and " + col + " = ?"
// 		}

// 		// cols = append(cols, col)
// 		value = append(value, val)
// 	}

// 	var count int64
// 	var err error
// 	if count, err = GetDB().Where(where, value...).Count(data); err != nil {
// 		logcus.Print("error", fmt.Sprintf("Get count data error from db. error: %s", err.Error()))
// 		return err
// 	} else if count > 0 {
// 		if _, err = GetDB().Where(where, value...).Update(data); err != nil {
// 			logcus.Print("error", fmt.Sprintf("Update data error to db. error: %s| data: %#v", err.Error(), data))
// 		}
// 	} else {
// 		if _, err = GetDB().Insert(data); err != nil {
// 			logcus.Print("error", fmt.Sprintf("Insert data error to db. data: %#v", data))
// 		}
// 	}

// 	// 要清除缓存的key
// 	if err == nil && len(keys) != 0 {
// 		if nums := cache.GetRedis().Del(keys...).Val(); int(nums) != len(keys) {
// 			logcus.Print("error", fmt.Sprintf("Clear the key from redis not complete. | plan: %d | now: %d | keys: %#v", len(keys), nums, keys))
// 		}
// 	}

// 	return err
// }

// QuerySQL 通用查询语句
func (bm *BaseModel) QuerySQL(sql string, value []interface{}, useCache bool, data interface{}, do func(data interface{})) error {
	cacheKey := ""
	if useCache {
		StartTime := time.Now()
		str := sql
		if value != nil {
			str = fmt.Sprint(sql, value)
		}
		cacheKey = fmt.Sprintf("%x", md5.Sum([]byte(str)))

		_, err := cache.GetCache(cacheKey, &data)
		if err == nil {
			bm.SetDebug(fmt.Sprintf("{Get data from local cache Key: %s ("+public.TimeCost(StartTime)+")}", cacheKey), 1)
		}

		return err
	}

	StartTime := time.Now()
	err := GetDB().Sql(sql, value...).Find(data)
	if err != nil {
		// fmt.Println(err.Error(), sql, value)
		logcus.Print("error", fmt.Sprintf("Get data from mysql error: %s | sql: %s | params: %#v", err.Error(), sql, value))
	} else {
		bm.SetDebug(fmt.Sprintf("{Get data from Mysql ("+public.TimeCost(StartTime)+"): %s}", data), 1)
		if cacheKey != "" {
			cache.SetCache(cacheKey, data, 600)
		}
	}

	return err
}
