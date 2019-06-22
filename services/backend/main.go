package main

import (
	"fmt"

	"github.com/houseofbosons/houseofbosons/services/backend/db"
)

func main() {
	_, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
