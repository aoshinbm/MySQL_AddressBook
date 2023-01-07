package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Contacts struct {
	Address_id int64
	Firstname  string
	Lastname   string
	Address    string
	City       string
	State      string
	Phoneno    int
	Email      string
}

var db *sql.DB

func main() {
	var err error
	//"UserName:Password@tcp(portNumber)/databaseName"
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/AddressBook")
	addres, err := listAll()
	if err != nil {
		fmt.Println("some error")
	} else {
		fmt.Println(addres)
	}
}
func listAll() ([]Contacts, error) {
	var contact []Contacts
	rows, err := db.Query("SELECT * FROM contact; ")
	if err != nil {
		return nil, fmt.Errorf("error in query all contact: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var addres Contacts
		if err := rows.Scan(&addres.Address_id, &addres.Firstname, &addres.Lastname,
			&addres.Address, &addres.City, &addres.State, &addres.Phoneno, &addres.Email); err != nil {
			return nil, fmt.Errorf("error in query all contact: %v",
				err)
		}
		contact = append(contact, addres)
	}
	return contact, nil
}
