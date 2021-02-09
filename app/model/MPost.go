package model

/*
 * @Author: Cedar
 * @Date: 2020-11-09 16:25:02
 * @LastEditors: Cedar
 * @LastEditTime: 2021-01-28 11:28:26
 * @FilePath: /Mango/app/model/Post.go
 */

// Post 表结构
type Posts struct {
	Id              int64  `json:"id"`
	AuthorId        int    `xorm:"author_id" json:"author_id"`
	CategoryId      int    `xorm:"category_id" json:"category_id"`
	Title           string `xorm:"title" json:"title"`
	SeoTitle        string `xorm:"seo_title" json:"seo_title"`
	Excerpt         string `xorm:"excerpt" json:"excerpt"`
	Body            string `xorm:"body" json:"body"`
	Image           string `xorm:"image" json:"image"`
	Slug            string `xorm:"slug" json:"slug"`
	MetaDescription string `xorm:"meta_description" json:"meta_description"`
	MetaKeywords    string `xorm:"meta_keywords" json:"meta_keywords"`
	Status          string `xorm:"status" json:"status"`
	Featured        string `xorm:"featured" json:"featured"`
	CreatedAt       string `xorm:"created_at" json:"-"`
	UpdatedAt       string `xorm:"updated_at" json:"-"`
}
