/*
 * @Author: Cedar
 * @Date: 2021-02-09 09:59:11
 * @LastEditors: Cedar
 * @LastEditTime: 2021-02-09 17:31:47
 * @FilePath: /Mango/app/model/MUsers.go
 */
package model

import (
	"time"
)

type Users struct {
	Id            int64     `json:"id"`
	GroupId       string    `xorm:"group_id" json:"group_id"`
	DeviceGroup   string    `xorm:"device_group" json:"device_group"`
	Username      string    `xorm:"username" json:"username"`
	Password      string    `xorm:"password" json:"-"`
	Nickname      string    `xorm:"nickname" json:"nickname"`
	Qq            string    `xorm:"qq" json:"qq"`
	Email         string    `xorm:"email" json:"email"`
	Phone         string    `xorm:"phone" json:"phone"`
	Status        string    `xorm:"status" json:"status"`
	Avatar        string    `xorm:"avatar" json:"avatar"`
	AvatarThumb   string    `xorm:"avatar_thumb" json:"avatar_thumb"`
	OpenidQq      string    `xorm:"openid_qq" json:"-"`
	OpenidWeixin  string    `xorm:"openid_weixin" json:"-"`
	Cert          int       `xorm:"cert" json:"cert"`
	CertReason    string    `xorm:"cert_reason" json:"cert_reason"`
	CertPic       string    `xorm:"cert_pic" json:"-"`
	RegIp         string    `xorm:"reg_ip" json:"-"`
	LoginIp       string    `xorm:"login_ip" json:"-"`
	LoginLocation string    `xorm:"login_location" json:"-"`
	DeviceId      string    `xorm:"device_id" json:"-"`
	LastLoginTime string    `xorm:"last_login_time" json:"-"`
	LastLoginIp   string    `xorm:"last_login_ip" json:"-"`
	LoginNum      int       `xorm:"login_num" json:"-"`
	CreatedId     int       `xorm:"created_id" json:"-"`
	CreatedAt     time.Time `xorm:"created_at" json:"-"`
	UpdatedId     int       `xorm:"updated_id" json:"-"`
	UpdatedAt     time.Time `xorm:"updated_at" json:"-"`
	DeleteId      int       `xorm:"deleted_id" json:"-"`
	DeletedAt     time.Time `xorm:"deleted_at" json:"-"`
}
