package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var (
	host     string
	port     int
	user     string
	password string
	dbname   string
)

func init() {
	host = os.Getenv("POSTGRES_HOST")

	var err error
	port, err = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		panic("invalid POSTGRES_PORT")
	}

	user = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname = os.Getenv("POSTGRES_DB_NAME")
}

// ConnectDB ...
func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return db, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}
