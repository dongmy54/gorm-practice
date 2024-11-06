package main

import (
	"encoding/json"
	"log"

	"gorm_practice/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 由于这里没有密码所以省略password=号后面没有值
	dsn := "host=localhost user=dongmingyan dbname=gorm_practice port=5432 password= sslmode=disable"

	// 打开数据库连接
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 检查数据库连接是否成功
	if err := db.Exec("SELECT 1").Error; err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("Successfully connected to the database!")

	// 自动迁移数据库模式 需要显示的写出来 否则不会迁移
	if err := db.AutoMigrate(models.User{}, models.CreditCard{}, models.Address{}, models.Student{}, models.Course{}, models.Product{}, models.Account{}); err != nil { // 导入 models 包中的 User 模型
		log.Fatalf("failed to auto migrate: %v", err)
	}

	// 创建一个切片
	// dataSlice := []string{"element1", "element2", "element3"}

	// // 序列化切片为JSON
	// jsonData, err := json.Marshal(dataSlice)
	// if err != nil {
	// 	panic(err)
	// }

	// db.Create(&models.User{
	// 	Name: "John Doe",
	// 	Data: jsonData,
	// })

	var user models.User
	db.Last(&user)
	var dataSlice []string
	if err := json.Unmarshal(user.Data, &dataSlice); err != nil {
		log.Printf("failed to unmarshal data: %v", err)
	}
	log.Printf("User: %#v", dataSlice)
}
