package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/voioc/mango/app/model"
)

/*
 * @Author: Cedar
 * @Date: 2021-01-28 11:08:27
 * @LastEditors: Cedar
 * @LastEditTime: 2021-02-09 17:37:08
 * @FilePath: /Mango/app/handler/Vod.go
 */

// VodIndex 获取首页数据流
func VodIndex(c *gin.Context) {

	bm := model.NewBaseModel(c)

	page := c.DefaultQuery("page", "1")
	data := bm.GetVodAuthorList(page)
	// extension := bm.GetExtensionList(etype)

	// result := map[string]interface{}{"operation": operation, "extension": extension}

	// data := []model.Vods{}
	// if err := model.GetDB().Where("status = ? ", 1).Limit(20).Find(&data); err != nil {
	// 	fmt.Println(err.Error())
	// 	// logcus.Print("error", "Get data from mysql error: "+err.Error())
	// } else {
	// 	// bm.SetDebug(fmt.Sprintf("{Get data from Mysql ("+public.TimeCost(StartTime)+"): %s}", data), 1)
	// 	// cache.SetCache(key, data, -1)
	// }

	c.JSON(http.StatusOK, bm.GetOutPut(data))
}
