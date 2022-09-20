package models

import (
	"github.com/go-crud/db"

	"github.com/go-crud/uteis"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetProducts() []Product {
	db := db.DbConnect()
	productsQuery, err := db.Query("select * from products")
	uteis.CheckIfExisteError(err)

	p := Product{}
	products := []Product{}

	for productsQuery.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productsQuery.Scan(&id, &name, &description, &price, &quantity)
		uteis.CheckIfExisteError(err)

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateProduct(name string, description string, price float64, quantity int) {
	db := db.DbConnect()

	insertQuery, err := db.Prepare("INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4)")
	uteis.CheckIfExisteError(err)

	insertQuery.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.DbConnect()

	deleteQuery, err := db.Prepare("DELETE FROM products WHERE id=$1")
	uteis.CheckIfExisteError(err)

	deleteQuery.Exec(id)
	defer db.Close()
}

func GetProduct(id string) Product {
	db := db.DbConnect()

	productQuery, err := db.Query("SELECT * FROM products WHERE id=$1", id)
	uteis.CheckIfExisteError(err)

	productToUpdate := Product{}

	for productQuery.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productQuery.Scan(&id, &name, &description, &price, &quantity)
		uteis.CheckIfExisteError(err)

		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Quantity = quantity
	}

	defer db.Close()
	return productToUpdate
}

func Update(id int, name string, description string, price float64, quantity int) {
	db := db.DbConnect()

	updateQuery, err := db.Prepare("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5")
	uteis.CheckIfExisteError(err)

	updateQuery.Exec(name, description, price, quantity, id)
	defer db.Close()
}
