package main

import (
	_ "github.com/lib/pq"
)

// delete all rows from the crypto_csv_raw_data table and insert the data from the csv file
func TransformRawData() ([]Transaction, error) {
	// Load the Postgres URI from environment variables or the .env file
	return []Transaction{}, nil
}
