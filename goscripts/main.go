package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	file, err := os.Open("transactions.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Load the Postgres URI from environment variables or the .env file
	postgresURI := os.Getenv("POSTGRES_URI")
	if postgresURI == "" {
		fmt.Println("Postgres URI not set in environment variables")
		file, err := os.Open("../.env")
		if err != nil {
			fmt.Println("Error opening .env file:", err)
			return
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			if strings.HasPrefix(scanner.Text(), "postgresql://") {
				postgresURI = scanner.Text()
				continue
			}
		}
		defer file.Close()
	}

	// Open the database connection
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Delete all rows from table
	_, err = db.Exec("DELETE FROM piggybank.crypto_csv_raw_data")
	if err != nil {
		fmt.Println("Error deleting rows:", err)
		return
	}

	total := 0
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fmt.Println(scanner.Text())

	// Build query for bulk insert
	bulkQuery := "INSERT INTO piggybank.crypto_csv_raw_data (id, \"timestamp_UTC\") VALUES"

	for true {
		ok := scanner.Scan()
		//if total > 5000 {
		//break
		//}
		if !ok {
			break
		}
		splitRow := strings.Split(scanner.Text(), ",")
		bulkQuery += fmt.Sprintf(" ('%d', '%s'),", total, splitRow[0])
		total += 1
		fmt.Println("Total lines:", total)
	}
	bulkQuery = bulkQuery[:len(bulkQuery)-1] + ";"
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	_, err = db.Exec(bulkQuery)
	if err != nil {
		fmt.Println("Error executing bulk insert:", err)
	}
}

//func addCsvRowToDatabase(db *sql.DB, row string) {
//	splitRow := strings.Split(row, ",")
//	if len(splitRow) != 11 {
//		fmt.Println("Invalid row:", row)
//		return
//	}
//	//timestamp, err := time.Parse(time.RFC3339, splitRow[0])
//	//if err != nil {
//	//	fmt.Println("Invalid timestamp:", splitRow[0])
//	//	return
//	//}
//	//timestamp := splitRow[0]
//	_, err := db.Exec("INSERT INTO piggybank.crypto_csv_raw_data (\"timestamp_UTC\") VALUES($1)", splitRow[0])
//	if err != nil {
//		fmt.Println("Error inserting row:", err)
//		return
//	}
//	fmt.Println(splitRow)
//}
