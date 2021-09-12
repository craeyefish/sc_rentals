package main

import (
	"context"
	"database/sql"
	"rentals/items"
	itemdb "rentals/items/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
)

var (
	// RentalsDB is the db used in the item package.
	RentalsDB *sql.DB
)

func main() {
	ctx := context.TODO()

	RentalsDB, err := sql.Open("mysql", "root:naic19..@tcp(127.0.0.1:3306)/rentals")
	if err != nil {
		panic(err.Error())
	}
	defer RentalsDB.Close()

	itemdb.Create(ctx, RentalsDB, items.Item{
		Name:  "TestItem",
		Type:  items.TypeChair,
		Value: decimal.NewFromInt(100),
	})

}
