package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goex"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//create
	// db.Create(&Product{
	// 	Name:  "Celular",
	// 	Price: 1800.12,
	// })
	//create batch
	products := []Product{
		{Name: "Notebook", Price: 1500.0},
		{Name: "Tablet", Price: 800.50},
		{Name: "Monitor", Price: 1000.0},
		{Name: "Mouse", Price: 50.0},
		{Name: "Teclado", Price: 120.0},
	}
	db.Create(&products)

}
