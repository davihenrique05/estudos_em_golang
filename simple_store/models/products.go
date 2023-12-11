package models

import (
	"fmt"

	"simple_store/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	products := []Product{}

	db := db.ConectWithDb()
	defer db.Close()

	selectProducts, err := db.Query("select * from products")

	if err != nil {
		fmt.Println(err)
	}

	for selectProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		p := Product{Id: id, Name: name, Description: description, Price: price, Quantity: quantity}

		products = append(products, p)
	}

	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.ConectWithDb()
	defer db.Close()

	insertData, err := db.Prepare("insert into products (name, description, price, quantity) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, price, quantity)
}

func DelectProduct(id string) {
	db := db.ConectWithDb()
	defer db.Close()

	deleteData, err := db.Prepare("delete from products where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteData.Exec(id)
}
