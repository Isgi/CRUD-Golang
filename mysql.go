package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB  {
	// db, err := sql.Open("mysql", "root:@/nduwur")
	db, err := sql.Open("mysql", "ipaddress:2wsx1qaz@tcp(192.168.3.39:3306)/ipaddress")
	if err != nil {
		log.Fatal(err.Error(), "Could not connect to database")
	}
	return db
}
// db, err := sql.Open("mysql", "root:@/nduwur")
// db, err := sql.Open("mysql", "ipaddress:2wsx1qaz@192.168.3.39:3306/ipaddress")
