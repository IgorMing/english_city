package database

import (
	"database/sql"
	"fmt"

	"github.com/IgorMing/englishcity/utils"
	_ "github.com/lib/pq"
)

const (
	user     = "igor"
	password = "Igor1993"
	dbname   = "englishcity"
)

var database *sql.DB

// Connect with the server
func Connect() error {
	psqlInfo := fmt.Sprintf("user=%s "+
		"password=%s dbname=%s sslmode=disable",
		user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return utils.HandleError("Error opening the connection.", err)
	}

	err = db.Ping()
	if err != nil {
		return utils.HandleError("Error trying a ping with the database", err)
	}

	database = db

	fmt.Println("Connected with success")
	return nil
}

func disconnect() {
	database.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
