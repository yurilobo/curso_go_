package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() { //cria coneção com banco de dados
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goex")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Notebook", 1800.12)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

}

//cira uma  função para inserir um novo dado

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO product (id, name, price) VALUES(?, ?,?)") //usamos ?
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil

}
