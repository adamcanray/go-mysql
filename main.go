package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Id   int    `json:"id"` // reflect, to inspect struct fields
	Name string `json:"name"`
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1)/go-mysql")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO product (name, price, stock) VALUES ('product1', 20000, 10)")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Yay, values added!")

	result, err := db.Query("SELECT id, name FROM product")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var product Product
		err = result.Scan(&product.Id, &product.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(product.Id, product.Name)
	}

}
