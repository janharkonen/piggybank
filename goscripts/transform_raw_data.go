package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type CurrencySet map[string]struct{}
type YearList []int

// transform the raw data from the database into the transaction struct
func TransformRawData(db *sql.DB) ([]Transaction, CurrencySet, YearList, error) {
	// Init variable
	var transaction Transaction
	var rawDataRow RawDataRow
	transactions := make([]Transaction, 0)
	currencySet := make(CurrencySet)
	yearList := make(YearList, 0)
	total := 0

	// Query database
	rows, err := db.Query("SELECT * FROM piggybank.crypto_csv_raw_data ORDER BY \"timestamp_UTC\" ASC")
	if err != nil {
		return []Transaction{}, nil, nil, err
	}
	defer rows.Close()

	// Loop through rows and fill transactions list
	for rows.Next() {
		total += 1
		fmt.Println("Total rows:", total)
		rawDataRow, err = convertFromDBtoRawDataRow(rows)
		if err != nil {
			return []Transaction{}, nil, nil, err
		}
		currencySet[rawDataRow.ToCurrency] = struct{}{}

		transaction, err = convertToTransaction(rawDataRow)
		if err == nil {
			transactions = append(transactions, transaction)
		}
	}

	// Construct list of years
	var firstYear int = transactions[0].Aikaleima.Year()
	var lastYear int = rawDataRow.TimestampUTC.Year()
	for year := firstYear; year <= lastYear; year++ {
		yearList = append(yearList, year)
	}

	return transactions, currencySet, yearList, nil
}

func convertFromDBtoRawDataRow(rows *sql.Rows) (RawDataRow, error) {
	var rawDataRow RawDataRow
	err := rows.Scan(
		&rawDataRow.TimestampUTC,
		&rawDataRow.TransactionDesc,
		&rawDataRow.Currency,
		&rawDataRow.Amount,
		&rawDataRow.ToCurrency,
		&rawDataRow.ToAmount,
		&rawDataRow.NativeCurrency,
		&rawDataRow.NativeAmountUSD,
		&rawDataRow.TransactionKind,
		&rawDataRow.TransactionHash,
		&rawDataRow.NativeAmount,
		&rawDataRow.id,
	)
	return rawDataRow, err
}

func convertToTransaction(rawDataRow RawDataRow) (Transaction, error) {
	// Init
	var transaction Transaction

	// Condition
	transaction.Aikaleima = rawDataRow.TimestampUTC
	switch rawDataRow.TransactionKind {
	case "crypto_purchase":
		transaction.Tyyppi = "BUY"
		transaction.Kryptovaluutta = rawDataRow.Currency
		transaction.HintaEUR = rawDataRow.NativeAmount
		transaction.MääräKryptovaluuttana = rawDataRow.Amount
		transaction.EURPerKryptovaluutta = rawDataRow.NativeAmount / rawDataRow.Amount
		transaction.KryptovaluuttaaJäljellä = rawDataRow.Amount
	case "crypto_transfer":
		return Transaction{}, errors.New("not an error per se, the volumes are just too tiny")
	default:
		return Transaction{}, errors.New("unknown transaction kind")
	}
	return transaction, nil
}
