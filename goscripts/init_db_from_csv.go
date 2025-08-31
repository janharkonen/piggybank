package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// delete all rows from the crypto_csv_raw_data table and insert the data from the csv file
func InitDbFromCsv(db *sql.DB, csvFilePath string) error {
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	// Delete all rows from table
	_, err = db.Exec("DELETE FROM piggybank.crypto_csv_raw_data")
	if err != nil {
		return err
	}

	total := 0
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = 11
	header, err := reader.Read()
	if err != nil {
		return err
	}
	_ = header

	// Build query for bulk insert
	bulkQuery := "INSERT INTO piggybank.crypto_csv_raw_data (id, \"timestamp_UTC\", transaction_desc, currency, amount, to_currency, to_amount, native_currency, native_amount, \"native_amount_USD\", transaction_kind, transaction_hash) VALUES"

	for true {
		record, err := reader.Read()
		//if total > 5000 {
		//break
		//}
		if err != nil {
			break
		}
		// descontruct the row into variables
		id := total
		timestamp := record[0]
		transactionDesc := record[1]
		currency := record[2]
		amount := record[3]
		toCurrency := record[4]
		toAmount := record[5]
		nativeCurrency := record[6]
		nativeAmount := record[7]
		nativeAmountUSD := record[8]
		transactionKind := record[9]
		transactionHash := record[10]
		partialQuery := fmt.Sprintf(" ('%d', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s'),", id, timestamp, transactionDesc, currency, amount, toCurrency, toAmount, nativeCurrency, nativeAmount, nativeAmountUSD, transactionKind, transactionHash)
		if toAmount == "" {
			partialQuery = fmt.Sprintf(" ('%d', '%s', '%s', '%s', '%s', '%s', NULL, '%s', '%s', '%s', '%s', '%s'),", id, timestamp, transactionDesc, currency, amount, toCurrency, nativeCurrency, nativeAmount, nativeAmountUSD, transactionKind, transactionHash)
		}
		bulkQuery += partialQuery
		total += 1
		fmt.Println("Total lines:", total)
	}
	bulkQuery = bulkQuery[:len(bulkQuery)-1] + ";"
	_, err = db.Exec(bulkQuery)
	if err != nil {
		return err
	}
	return nil
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
