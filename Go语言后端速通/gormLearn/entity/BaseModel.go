package entity

import "time"

// BaseModel 所有属性首字母不大写会被gorm忽略
type BaseModel struct {
	ID         int        `gorm:"primarykey"`
	CreateTime *time.Time `gorm:"autoCreateTime;"`
	DeleteTime *time.Time `gorm:"autoCreateTime;"`
	// time.Time 总为一个有效值,数据库的对应字段不能为空 Not Null
	// *time.Time 则可以在数据库对应字段为空时使用
}
