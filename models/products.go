package models

import "github.com/andre2l2/go_mvc/db"

type Products struct {
	Id int
	Name string
	Description string
	Price float64
	Total int
}

func GetAllProducs() []Products {
	db := db.ConectToDatabase()
	selectAll, err := db.Query("select * from products order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Products{}
	products := []Products{}

	for selectAll.Next() {
		var id, total int
		var name, description string
		var price float64

		err = selectAll.Scan(&id, &name, &description, &price, &total)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Total = total

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func GetOneProduct(id string) Products {
	db := db.ConectToDatabase()
	getOne, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	
	products := Products{}

	for getOne.Next() {
		var id, total int
		var name, description string
		var price float64

		err := getOne.Scan(&id, &name, &description, &price, &total)
		if err != nil {
			panic(err.Error())
		}

		products.Id = id
		products.Name = name
		products.Description = description
		products.Price = price
		products.Total = total
	}

	defer db.Close()
	return products
}

func CreatNewProduct(name string, description string, price float64, total int) {
	db := db.ConectToDatabase()
	insertIntoDb, err := db.Prepare("insert into products (name, description, price, total) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertIntoDb.Exec(name, description, price, total)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConectToDatabase()
	deleteProduct, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}

func UpdateProduct(id string, name string, description string, price float64, total int) {
	db := db.ConectToDatabase()
	update, err := db.Prepare("update products set name=$1, description=$2, price=$3, total=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	update.Exec(name, description, price, total, id)
	defer db.Close()
}