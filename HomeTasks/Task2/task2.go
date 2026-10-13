package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "host=localhost port=5432 user=postgres password=js0770 dbname=myhomeworkdb sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	s := db.Stats()
	fmt.Println("Total open connections:", s.OpenConnections)
	fmt.Println("Active connections (in use):", s.InUse)
	fmt.Println("Free connections (idle, waiting in pool):", s.Idle)
}
