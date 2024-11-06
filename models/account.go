package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	// 组合索引 同时命名了idx_space_location 它用于两个字段上
	SpaceId    uint `gorm:"index:idx_space_location;not null"` // 组合索引
	LocationId uint `gorm:"index:idx_space_location;not null"` // 组合索引

	Age uint `gorm:"not null"` // 整型

	Name     string `gorm:"type:varchar(255);not null"`        // 字符串
	Email    string `gorm:"type:varchar(255);not null;index"`  // 字符串(普通索引)
	PhoneNum string `gorm:"type:varchar(255);not null;unique"` // 字符串（唯一索引）
	// 描述
	Description string `gorm:"type:text;not null"` // 文本类型
	// 余额
	Balance float64 `gorm:"not null"` // 浮点类型
	// 是否有效 默认有效
	Active bool `gorm:"default:true"` // 布尔类型 默认值true
	// 有效期
	ExpiredAt *time.Time // 时间类型
}
