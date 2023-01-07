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
	newContact := Contacts{
		Firstname: "Manu",
		Lastname:  "Frown",
		Address:   "PorshStreet",
		City:      "Mumbai",
		State:     "Maharashtra",
		Phoneno:   9876789332,
		Email:     "fro@gmail.com",
	}
	lastId, err := addContact(newContact)
	if err == nil {
		fmt.Println("Last insert id:", lastId)
	}
}
func addContact(con Contacts) (int64, error) {
	result, err := db.Exec("INSERT INTO contact (address_id,firstname,lastname,address,city,state,phoneno,email)VALUES (?, ?, ?,?,?,?,?,?)",
		con.Address_id, con.Firstname, con.Lastname, con.Address, con.City, con.State, con.Phoneno, con.Email)
	if err != nil {
		return 0, fmt.Errorf("add contact: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("add contact: %v", err)
	}
	return id, nil
}
