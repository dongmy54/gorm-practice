package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// User 模型定义
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Age      int    `gorm:"type:int"`
	State    string `gorm:"type:varchar(100)"`
	PhoneNum string `gorm:"type:varchar(100)"`

	CreditCard CreditCard // user has_one credit_card
	// 使用自定义的DataJSONB类型来存储json数据
	Hobbies   DataJSONB      `gorm:"type:jsonb;"`
	DeletedAt sql.NullTime   `gorm:"column:deleted_at"`
	Data      datatypes.JSON `gorm:"type:jsonb"` // 使用`datatypes.JSON`类型以使用jsonb
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

// 批量插入
// 其中某一个发生错误将导致本次整个插入失败
func BatchCreate(db *gorm.DB, users []*User, batch_size int) error {
	return db.CreateInBatches(users, batch_size).Error
}

// 这种方式批量插入是可以的，当发生冲突自动忽略
func BatchCreate1(db *gorm.DB, users []*User) error {
	if len(users) == 0 {
		return nil
	}

	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&users).Error
}

// 对数据批量操作
func BatchOperation(db *gorm.DB) error {
	results := make([]*User, 0)
	res := db.Model(User{}).Unscoped().FindInBatches(&results, 2, func(tx *gorm.DB, batch int) error {
		npUsers := make([]*User, 0)
		// 循环遍历数据
		for _, user := range results {
			log.Printf("user: %#v=====\n", user)
			if user.PhoneNum == "" {
				num := fmt.Sprintf("phoneNum_%d", user.ID)
				npUsers = append(npUsers, &User{Name: "空手机号用户", PhoneNum: num})
			}
		}

		// 找出来后 批量插入
		BatchCreate1(db, npUsers)
		return nil
	})

	return res.Error
}
