package models

import (
	base "gopher_store/Base"
)

type Products struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func SearchAllProducts() []Products {
	db := base.ConnectionForDataBase()

	selectAllProducts, err := db.Query("SELECT id, name, description, price, quantity FROM products")
	if err != nil {
		panic(err.Error())
	}

	p := Products{}
	products := []Products{}

	for selectAllProducts.Next() {
		var id int
		var name, description string
		var price float64
		var quantity int

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.ID = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func NewProducts(name, description string, price float64, quantity int) {
	db := base.ConnectionForDataBase()
	defer db.Close()

	insertProduct, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(name, description, price, quantity)
	defer insertProduct.Close()

}

func UpdateProducts(id string, name, description string, price float64, quantity int) {
	db := base.ConnectionForDataBase()
	defer db.Close()

	updateProduct, err := db.Prepare("UPDATE products SET name = $1, description = $2, price = $3, quantity = $4 WHERE id = $5")
	if err != nil {
		panic(err.Error())
	}
	defer updateProduct.Close()

	updateProduct.Exec(name, description, price, quantity, id)
}

func DeleteProducts(id int) {
	db := base.ConnectionForDataBase()
	defer db.Close()

	deleteProduct, err := db.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}
	defer deleteProduct.Close()

	deleteProduct.Exec(id)
}

func EditProducts(id string) Products {
	db := base.ConnectionForDataBase()
	SelectProduct, err := db.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	productToEdit := Products{}

	for SelectProduct.Next() {
		var id int
		var name, description string
		var price float64
		var quantity int

		err = SelectProduct.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		productToEdit.ID = id
		productToEdit.Name = name
		productToEdit.Description = description
		productToEdit.Price = price
		productToEdit.Quantity = quantity
	}
	defer db.Close()
	return productToEdit
}
