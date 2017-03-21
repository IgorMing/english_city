package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/IgorMing/englishcity/utils"
	_ "github.com/lib/pq"
)

var database *sql.DB

// Connect with the server
func Connect() error {
	host, port, user, password, envDB := os.Getenv("DB_PORT_5432_TCP_ADDR"), os.Getenv("DB_PORT_5432_TCP_PORT"), "postgres", os.Getenv("DB_ENV_POSTGRES_PASSWORD"), os.Getenv("DB_ENV_POSTGRES_DB")
	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "5432"
	}

	if user == "" {
		user = "postgres"
	}

	if password == "" {
		password = "Igor1993"
	}

	if envDB == "" {
		envDB = "english_city"
	}

	psqlInfo := fmt.Sprintf("user=%s "+
		"password=%s dbname=%s sslmode=disable host=%s port=%s",
		user, password, envDB, host, port)
	fmt.Printf(psqlInfo)

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
