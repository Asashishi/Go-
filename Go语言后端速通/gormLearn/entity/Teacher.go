package entity

import (
	"time"
)

type Teacher struct {
	BaseModel
	Name   string     `gorm:"type:varchar(32);"`
	Tno    int        `gorm:"unique"`
	Pwd    string     `gorm:"type:varchar(32);not null"`
	Tel    string     `gorm:"type:varchar(11);"`
	Birth  *time.Time `gorm:"type:datetime;"`
	Remark string     `gorm:"type:varchar(255);"`
}
