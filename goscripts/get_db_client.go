package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

// delete all rows from the crypto_csv_raw_data table and insert the data from the csv file
func GetDbClient() (*sql.DB, error) {
	// Load the Postgres URI from environment variables or the .env file
	postgresURI := os.Getenv("POSTGRES_URI")
	if postgresURI == "" {
		fmt.Println("Postgres URI not set in environment variables")
		envFile, err := os.Open("../.env")
		if err != nil {
			fmt.Println("Error opening .env file:", err)
			return nil, err
		}
		scanner := bufio.NewScanner(envFile)
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), "postgresql://") {
				postgresURI = scanner.Text()
				continue
			}
		}
		defer envFile.Close()
	}

	// Open the database connection
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return nil, err
	}

	return db, nil
}
