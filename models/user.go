package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// User 模型定义
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Age      int    `gorm:"type:int"`
	State    string `gorm:"type:varchar(100)"`

	CreditCard CreditCard // user has_one credit_card
	// 使用自定义的DataJSONB类型来存储json数据
	Hobbies DataJSONB      `gorm:"type:jsonb;"`
	Data    datatypes.JSON `gorm:"type:jsonb"` // 使用`datatypes.JSON`类型以使用jsonb
}

// 注意默认情况下gorm会同时帮我们创建id、created_at、updated_at、deleted_at
// 这三个字段

// Define a new type
type DataJSONB []string

// Implement the Valuer and Scanner interfaces
func (dj DataJSONB) Value() (driver.Value, error) {
	return json.Marshal(dj)
}

func (dj *DataJSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("[]byte assertion failed")
	}

	return json.Unmarshal(b, dj)
}

// scope方法
// 查询年龄大于xx的用户
func AgeGreaterThan(age int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("age > ?", age)
	}
}

// state为有效的用户
func ValidState(db *gorm.DB) *gorm.DB {
	return db.Where("state = ?", "valid")
}
