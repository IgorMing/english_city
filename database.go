package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	user     = "igor"
	password = "Igor1993"
	dbname   = "english_city"
)

var database *sql.DB

func connect() error {
	psqlInfo := fmt.Sprintf("user=%s "+
		"password=%s dbname=%s sslmode=disable",
		user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	database = db
	fmt.Println("Connected with success")
	return nil
}

func disconnect() {
	database.Close()
}

func makeSelect(query string) {
	rows, err := database.Query("SELECT * FROM room")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var room Room
		err = rows.Scan(&room.ID, &room.Name)
		if err != nil {
			panic(err)
		}

		fmt.Println("id | name")
		fmt.Printf("%3v | %8v\n", room.ID, room.Name)
	}
}

// 	fmt.Println("# Inserting values")
//
// 	var lastInsertId int
// 	err = db.QueryRow("INSERT INTO room(name) VALUES($1) returning id;", "testando!!!").Scan(&lastInsertId)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("last inserted id =", lastInsertId)
// }
