package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// setting up the connection
func main() {
	var err error
	//"UserName:Password@tcp(portNumber)/databaseName"
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/AddressBook")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
