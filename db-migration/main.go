package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	dbConn, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}

	dbConn.SetConnMaxLifetime(time.Minute * 3)
	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(10)

	if err := dbConn.Ping(); err != nil {
		panic(err)
	}

	// perform migration
	var (
		direction = "up"
		step      = 1
	)
	migrations := &migrate.FileMigrationSource{Dir: "./migrations"}
	var n int // "n" berapa migration yg di-applied
	if direction == "down" {
		n, err = migrate.ExecMax(dbConn, "mysql", migrations, migrate.Down, step)
	} else {
		n, err = migrate.ExecMax(dbConn, "mysql", migrations, migrate.Up, step)
	}
	if err != nil {
		log.Fatalf("failed to run migration, error: %s", err.Error())
	}

	log.Printf("successfully applied %d migration(s)", n)
}
