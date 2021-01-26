/*
 * @Author: Cedar
 * @Date: 2019-11-27 14:08:56
 * @LastEditors: Cedar
 * @LastEditTime: 2021-01-26 17:14:50
 * @FilePath: /Mango/app/model/MysqlModel.go
 */
package model

import (
	"log"
	"os"
	"sync"

	"github.com/voioc/coco/config"
	"github.com/voioc/coco/logcus"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var eg *xorm.EngineGroup
var lockMysql sync.Mutex

// Init 初始化连接
func Init() {
	conn()
}

// GetDB 获取数据库
func GetDB() *xorm.EngineGroup {
	if eg == nil {
		conn()
	}

	return eg
}

// GetMySQL mysql实例
func conn() {
	if eg == nil {
		lockMysql.Lock()
		defer lockMysql.Unlock()

		driver := config.GetConfig().GetString("db.driver")
		nodes := config.GetConfig().GetStringSlice("db.nodes")

		if driver != "" && nodes[0] != "" {
			// conns := []string{
			// 	"lej_admin:Y1Gp/bKwcBk6JSiz@tcp(m3112i.tjtn.db.lecloud.com:3112)/lej?charset=utf8mb4;",
			// 	"lej_admin:Y1Gp/bKwcBk6JSiz@tcp(s3112i.tjtn.db.lecloud.com:3112)/lej?charset=utf8mb4;",
			// }

			var err error
			eg, err = xorm.NewEngineGroup(driver, nodes)
			if err != nil {
				log.Fatalln("Can't connect the database:", err.Error())
			}

			sqlLog := config.GetConfig().GetString("log.sql")
			if sqlLog == "" {
				sqlLog = "/tmp/sql.log"
			}

			if sl, err := os.OpenFile(sqlLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err != nil {
				logcus.Print("info", "打开错误日志文件失败:"+err.Error())
			} else {
				// eg.SetLogger(xorm.NewSimpleLogger(logcus.GetLogger().Out))
				eg.SetLogger(xorm.NewSimpleLogger(sl))
			}

			eg.ShowSQL(true)
			eg.SetMaxIdleConns(10)
			eg.SetMaxOpenConns(100)
			eg.Logger().SetLevel(core.LOG_DEBUG)
		}
	}
}
