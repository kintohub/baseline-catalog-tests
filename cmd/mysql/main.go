package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func main() {
	host := os.Getenv("TARGET_HOST")
	mysqlConnectionStr := fmt.Sprintf("user:password@%v/dbname", host)
	log.Printf("Attempting to connect to %v", mysqlConnectionStr)

	_, err := sql.Open("mysql", mysqlConnectionStr)

	if err != nil {
		panic(err)
	}

	log.Printf("Successfully ran mysql test")
}
