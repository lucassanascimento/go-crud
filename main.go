package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("template/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

	products := getProducts()
	fmt.Println(products)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := getProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func dbConnect() *sql.DB {
	connection := "user=postgres dbname=postgres password=tasken host=localhost port=5433 sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func getProducts() []Product {
	db := dbConnect()
	productsQuery, err := db.Query("select * from products")
	checkIfIsError(err)

	p := Product{}
	products := []Product{}

	for productsQuery.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productsQuery.Scan(&id, &name, &description, &price, &quantity)
		checkIfIsError(err)

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
		db.Close()
	}
	return products
}

func checkIfIsError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
