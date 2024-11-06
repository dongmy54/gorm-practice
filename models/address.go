package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	// 省份
	Province string `gorm:"type:varchar(20);not null"`
	// 城市
	City string `gorm:"type:varchar(20);not null"`
	// 区
	District string `gorm:"type:varchar(20);not null"`
	// 详细地址
	Detail string `gorm:"type:varchar(100);not null"`
	UserId uint
	User   *User // 属于user 避免循环引用
}
