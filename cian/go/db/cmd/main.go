package main

import "database/sql"

func main() {
	db, err := sql.Open("mysql", "root:naic19..@tcp(127.0.0.1:3306)/rentals")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	result, err := db.Query("select * from items")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer result.Close()

}
