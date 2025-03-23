package main

import (
	"fmt"

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
	// products := []Product{
	// 	{Name: "Notebook", Price: 1500.0},
	// 	{Name: "Tablet", Price: 800.50},
	// 	{Name: "Monitor", Price: 1000.0},
	// 	{Name: "Mouse", Price: 50.0},
	// 	{Name: "Teclado", Price: 120.0},
	// }
	// db.Create(&products)

	//select one
	// var product Product
	// db.First(&product, 2)
	// db.First(&product, "name = 7", "Mouse")
	// fmt.Println(product)

	//select all
	var products []Product
	db.Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

	// //update
	// db.Model(&product).Update("Price", 1600.0)

	// //delete
	// db.Delete(&product)
}
