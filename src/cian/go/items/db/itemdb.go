package itemdb

import (
	"context"
	"database/sql"
	"log"
	"rentals/items"
	"time"
)

// Create a new item in the database, fails if exists.
func Create(ctx context.Context, dbc *sql.DB, item items.Item) error {
	_, err := dbc.ExecContext(
		ctx,
		"insert into items (name, type, value, init_date) values(?, ?, ?, ?)",
		item.Name, item.Type, item.Value, time.Now(),
	)

	return err
}

// List returns all the items.
func List(ctx context.Context, dbc *sql.DB) ([]items.Item, error) {
	s, err := dbc.QueryContext(
		ctx,
		"select * from items",
	)
	if err != nil {
		log.Print(ctx, err)
		return nil, err
	}

	var res []items.Item

	for s.Next() {
		var i items.Item

		err = s.Scan(&i.Name, &i.Code, &i.Type, &i.Value, &i.InitDate)
		if err != nil {
			return nil, err
		}

		res = append(res, i)
	}

	return res, nil
}

// Updat all the name, type and value fields of an item.
func Update(ctx context.Context, dbc *sql.DB, item items.Item) error {
	_, err := dbc.ExecContext(
		ctx,
		"update items set name=?, type=?, value=?",
		item.Name, item.Type, item.Value,
	)

	return err
}
