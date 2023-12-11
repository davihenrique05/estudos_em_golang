package models

import (
	"fmt"

	"storeapp/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetProductById(id string) Product {
	db := db.ConectWithDb()
	defer db.Close()

	selectProduct, err := db.Query("select * from products where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for selectProduct.Next() {
		var id, quantitiy int
		var name, description string
		var price float64

		err := selectProduct.Scan(&id, &name, &description, &price, &quantitiy)

		if err != nil {
			panic(err.Error())
		}

		product = *ConvertDataToProduct(id, name, description, price, quantitiy)
	}

	return product
}

func GetAllProducts() []Product {
	products := []Product{}

	db := db.ConectWithDb()
	defer db.Close()

	selectProducts, err := db.Query("select * from products order by id asc")

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

func UpdateProduct(product Product) {
	db := db.ConectWithDb()
	defer db.Close()

	updateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(product.Name, product.Description, product.Price, product.Quantity, product.Id)
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

func ConvertDataToProduct(id int, name, description string, price float64, quantity int) *Product {
	product := Product{id, name, description, price, quantity}
	return &product
}
