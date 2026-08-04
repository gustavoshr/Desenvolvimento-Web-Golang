package models

import (
	"gopher_store/Base"
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
