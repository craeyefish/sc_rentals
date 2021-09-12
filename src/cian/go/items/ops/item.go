package item

import (
	"context"
	"database/sql"
	"rentals/items"
	itemdb "rentals/items/db"

	_ "github.com/go-sql-driver/mysql"
)

// CreateItem adds an item to the items table in the rentals db.
func CreateItem(ctx context.Context, item items.Item) error {
	db, err := sql.Open("mysql", "root:naic19..@tcp(127.0.0.1:3306)/rentals")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return itemdb.Create(ctx, db, item)
}
