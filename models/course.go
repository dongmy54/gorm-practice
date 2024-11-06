package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(100);not null"`
	Students []Student `gorm:"many2many:student_courses;"` // 学生与课程多对多关系
}
