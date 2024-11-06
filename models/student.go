package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100);not null"`
	Age     int8
	Courses []Course `gorm:"many2many:student_courses;"` // 定义多对多关系 一个学生有门课程

	// PS：不需要显示的定义student_coures结构体 它会自动创建表student_courses

	// 普通索引直接 `gorm:"index"`就行
	No string `gorm:"type:varchar(100);uniqueIndex"` // 创建唯一索引
}

func (s *Student) BeforeSave(tx *gorm.DB) (err error) {
	//如果学号是空字符串 或者不存在
	if s.No == "" {
		// 生成一个随机的10位随机字符串做学号
		s.No = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	return
}
