package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	ID    int64
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

	// customers, err := findAllCustomers(dbConn)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(customers)

	// customer, err := findByID(dbConn, 2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(customer)

	newCustomer, err := addCustomer(dbConn, Customer{
		Name:  "Andrew",
		Email: "andrew@yahoo.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newCustomer)
}

func findAllCustomers(dbConn *sql.DB) ([]Customer, error) {
	rows, err := dbConn.Query("SELECT id, name, email FROM customer")
	if err != nil {
		return nil, err
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

	return customers, nil
}

func findByID(dbConn *sql.DB, id int64) (*Customer, error) {
	row := dbConn.QueryRow("SELECT id, name, email FROM customer WHERE id=?", id)
	var customer Customer
	err := row.Scan(&customer.ID, &customer.Name, &customer.Email)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func addCustomer(dbConn *sql.DB, data Customer) (*Customer, error) {
	result, err := dbConn.Exec("INSERT INTO customer (name, email, created_at) VALUES (?, ?, ?)",
		data.Name,
		data.Email,
		time.Now().UTC())
	if err != nil {
		return nil, err
	}

	insertedID, _ := result.LastInsertId()
	data.ID = insertedID

	return &data, nil
}
