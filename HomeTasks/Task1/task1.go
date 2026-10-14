package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "host=localhost port=5432 user=postgres password=secretpaswword dbname=myhomeworkdb sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection successfully.")
}
