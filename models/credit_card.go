package models

import "gorm.io/gorm"

type CreditCard struct {
	gorm.Model
	Number string // 不指定具体长度会退化为text
	UserId uint   //

	// 这里用*User主要作用是避免循环引用（User中已经有CreditCard了）
	User *User // 使其可以通过credit_card.User
}
