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

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	db.Close()
	return products
}
