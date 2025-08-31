package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type RawDataRow struct {
	id              string
	TimestampUTC    string
	TransactionDesc string
	Currency        string
	Amount          float64
	ToCurrency      string
	ToAmount        float64
	NativeCurrency  string
	NativeAmount    float64
	NativeAmountUSD float64
	TransactionKind string
	TransactionHash string
}

// transform the raw data from the database into the transaction struct
func TransformRawData(db *sql.DB) ([]Transaction, error) {
	rows, err := db.Query("SELECT * FROM piggybank.crypto_csv_raw_data")
	transactions := []Transaction{}
	total := 0
	if err != nil {
		return []Transaction{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var rawDataRow RawDataRow
		total += 1
		fmt.Println("Total rows:", total)
		var transaction Transaction
		err := rows.Scan(
			&rawDataRow.id,
			&rawDataRow.TimestampUTC,
			&rawDataRow.TransactionDesc,
			&rawDataRow.Currency,
			&rawDataRow.Amount,
			&rawDataRow.ToCurrency,
			&rawDataRow.ToAmount,
			&rawDataRow.NativeCurrency,
			&rawDataRow.NativeAmount,
			&rawDataRow.NativeAmountUSD,
			&rawDataRow.TransactionKind,
			&rawDataRow.TransactionHash,
		)
		if err != nil {
			return []Transaction{}, err
		}
		transactions = append(transactions, transaction)
		//if err != nil {
		//	return []Transaction{}, err
		//}
		//transactions = append(transactions, transaction)
	}

	return transactions, nil
}
