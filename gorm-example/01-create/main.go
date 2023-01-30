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

func main() {
	dsn := "root:123456@tcp(localhost:3306)/example?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}

	m := db.Migrator()
	m.CreateTable(&Product{})

	// create a data
	p := &Product{Code: "D01", Price: 100}
	res := db.Create(p)
	fmt.Println(p.ID)
	if res.Error != nil {
		fmt.Println(res.Error)
	}

	// create some data
	products := []*Product{
		{Code: "D02", Price: 20},
		{Code: "D03", Price: 50},
		{Code: "D04", Price: 75}}
	res = db.Create(products)
	for _, item := range products {
		fmt.Println(item.ID)
	}
}
