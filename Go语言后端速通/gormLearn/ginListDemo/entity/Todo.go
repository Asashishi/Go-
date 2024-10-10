package entity

import "gorm.io/gorm"

type Todo struct {
	// gorm.Model带一个软删除
	gorm.Model
	// 代办事项标题
	Title string `form:"title" json:"title" gorm:"unique"`
	// 是否完成
	Status int `form:"status" json:"status"`
}
