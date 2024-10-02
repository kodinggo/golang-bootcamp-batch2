package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	ID    int
	Name  string
	Email string
}

func main() {
	dbConn, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}

	dbConn.SetConnMaxLifetime(time.Minute * 3)
	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(10)

	if err := dbConn.Ping(); err != nil {
		panic(err)
	}

	rows, err := dbConn.Query("SELECT id, name, email FROM customer")
	if err != nil {
		log.Fatal(err)
	}
	var customers []Customer
	for rows.Next() {
		var customer Customer
		err = rows.Scan(&customer.ID,
			&customer.Name,
			&customer.Email)
		if err != nil {
			log.Println("scan customer data:", err.Error())
			continue
		}
		customers = append(customers, customer)
	}

	fmt.Println(customers)
}
