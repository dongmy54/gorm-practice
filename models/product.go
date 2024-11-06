package models

import (
	"errors"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	StockNum    int // 库存数量
}

// 库存扣减
func DeductStock(db *gorm.DB, productID uint, quantity int) error {
	var product Product

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id =?", productID).First(&product).Error; err != nil {
			log.Printf("查询商品失败: %v\n", err)
			return err
		}

		newStockNum := product.StockNum - quantity
		if newStockNum < 0 {
			log.Println("库存不足")
			return errors.New("库存不足")
		}

		err := tx.Model(&product).Update("stock_num", newStockNum).Error
		if err != nil {
			log.Printf("更新库存失败: %v\n", err)
			return err
		}

		log.Println("库存扣减成功")
		return nil
	})

	return err
}

func CreateProduct(db *gorm.DB, name string, description string, price float64, stockNum int) error {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
		StockNum:    stockNum,
	}

	if err := db.Save(&product).Error; err != nil {
		log.Printf("创建商品失败: %v", err)
		return err
	}
	return nil
}
