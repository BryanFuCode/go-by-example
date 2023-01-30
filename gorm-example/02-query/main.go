package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint   `gorm:"primarykey"`
	Code  string `gorm:"column:code"`
	Price uint   `gorm:"column:price"`
}

func (Product) TableName() string {
	return "product"
}

func GetProduct(p Product) {
	fmt.Printf("ID : %d\n", p.ID)
	fmt.Printf("Code : %s\n", p.Code)
	fmt.Printf("Price : %d\n", p.Price)
}

func main() {
	dsn := "root:123456@tcp(localhost:3306)/example?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}

	p := &Product{}

	// 根据主键查询第一条记录
	// 查询不到数据会返回ErrRecordNotFound
	// SELECT * FROM product ORDER BY id LIMIT 1;
	db.First(p)
	GetProduct(*p)

	// 随机获取一条记录
	// SELECT * FROM products LIMIT 1;
	db.Take(p)
	GetProduct(*p)

}
