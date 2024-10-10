package entity

import "time"

type Student struct {
	BaseModel
	Name   string `gorm:"type:varchar(32)"`
	Sno    int    `gorm:"unique"`
	Pwd    string `gorm:"type:varchar(100);not null"`
	Tel    string `gorm:"type:varchar(11);"`
	Gender int
	Birth  *time.Time `gorm:"type:datetime;"`
	Remark string     `gorm:"type:varchar(255);"`

	// 多对一
	ClassID int
	// 映射外键结构体
	Class Class `gorm:"foreignkey:ClassID;constraint:OnDelete:CASCADE;"`

	// 多对多映射及联级操作 gorm自动创建关联表 可以访问对应结构体对象
	Courses []Course `gorm:"many2many:student2course;constraint:OnDelete:CASCADE;"`
	// 相当于
	/*
		type struct Student2Course{
			BaseModel
			ClassID int
			Class Class `gorm:"foreignkey:ClassID"`
			StudentID int
			Student Student "foreignkey:ClassID;constraint:OnDelete:CASCADE;"`
		}
	*/
}
