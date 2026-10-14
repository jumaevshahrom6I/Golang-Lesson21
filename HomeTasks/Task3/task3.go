package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var in = bufio.NewScanner(os.Stdin)

func ask(p string) string {
	fmt.Print(p)
	in.Scan()
	return strings.TrimSpace(in.Text())
}

func askInt(p string) int {
	n, _ := strconv.Atoi(ask(p))
	return n
}

func listBooks(db *sql.DB) {
	rows, err := db.Query("SELECT id, title, author, price, pages FROM books ORDER BY id")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id, price, pages int
		var title, author string
		if err := rows.Scan(&id, &title, &author, &price, &pages); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("%d | %s | %s | %d | %d\n", id, title, author, price, pages)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}

func findBook(db *sql.DB) {
	id := askInt("ID: ")
	var title, author string
	var price, pages int
	err := db.QueryRow("SELECT title, author, price, pages FROM books WHERE id = $1", id).
		Scan(&title, &author, &price, &pages)
	if err == sql.ErrNoRows {
		fmt.Println("Book not found")
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%d | %s | %s | %d | %d\n", id, title, author, price, pages)
}

func addBook(db *sql.DB) {
	title := ask("Title: ")
	author := ask("Author: ")
	price := askInt("Price: ")
	pages := askInt("Pages: ")
	var id int
	err := db.QueryRow("INSERT INTO books (title, author, price, pages) VALUES ($1, $2, $3, $4) RETURNING id",
		title, author, price, pages).Scan(&id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Added, id =", id)
}

func updateBook(db *sql.DB) {
	id := askInt("ID: ")
	title := ask("New title: ")
	author := ask("New author: ")
	price := askInt("New price: ")
	pages := askInt("New pages: ")
	res, err := db.Exec("UPDATE books SET title=$1, author=$2, price=$3, pages=$4 WHERE id=$5",
		title, author, price, pages, id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		fmt.Println("Book not found")
		return
	}
	fmt.Println("Updated")
}

func deleteBook(db *sql.DB) {
	id := askInt("ID: ")
	res, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		fmt.Println("Book not found")
		return
	}
	fmt.Println("Deleted")
}

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

	for {
		fmt.Println("\n1. Show all books\n2. Find book by ID\n3. Add book\n4. Update book\n5. Delete book\n0. Exit")
		switch ask("Choice: ") {
		case "1":
			listBooks(db)
		case "2":
			findBook(db)
		case "3":
			addBook(db)
		case "4":
			updateBook(db)
		case "5":
			deleteBook(db)
		case "0":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
