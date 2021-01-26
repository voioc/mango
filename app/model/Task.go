package model

/*
 * @Author: Cedar
 * @Date: 2020-11-09 16:25:02
 * @LastEditors: Cedar
 * @LastEditTime: 2020-11-10 10:36:41
 * @FilePath: /LeView/app/model/Task.go
 */

// Task 表结构
type Task struct {
	Id        int64  `json:"id"`
	Title     string `xorm:"title" json:"title"`
	SubTitle  string `xorm:"sub_title" json:"sub_title"`
	Type      int    `xorm:"type" json:"type"`
	Coin      int    `xorm:"coin" json:"coin"`
	Frequency int    `xorm:"frequency" json:"frequency"`
	Icon      string `xorm:"icon" json:"icon"`
	Url       string `xorm:"url" json:"url"`
	Button    string `xorm:"button" json:"button"`
	Remark    string `xorm:"remark" json:"remark"`
	Status    int    `xorm:"status" json:"status"`
	CreatedAt string `xorm:"created_at" json:"-"`
	UpdatedAt string `xorm:"updated_at" json:"-"`
	UpdatedId string `xorm:"updated_id" json:"-"`
	DeletedAt string `xorm:"updated_at" json:"-"`
}
