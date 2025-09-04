package main

import (
	"database/sql"
	"fmt"
	"time"
)

// transform the raw data from the database into the transaction struct
func WriteTransactionsToDB(db *sql.DB, transactions []Transaction, currencySet CurrencySet, yearList YearList) error {

	var laskettuOstohinta string
	var voitto string
	var jaljella_krypto string

	// Delete all rows from table
	_, err := db.Exec("DELETE FROM piggybank.crypto_refined")
	if err != nil {
		return err
	}

	bulkQuery := "INSERT INTO piggybank.crypto_refined (id, kryptovaluutta,\"aikaleima_UTC\", tyyppi, hinta_eur, maara_krypto, jaljella_krypto, laskettu_ostohinta, voitto, kommentti, eur_per_krypto) VALUES"
	total := -1
	partialQuery := ""
	for _, transaction := range transactions {
		total += 1
		fmt.Println("Total rows:", total)
		jaljella_krypto = "NULL"
		if transaction.KryptovaluuttaaJäljellä.Valid {
			jaljella_krypto = fmt.Sprintf("'%f'", transaction.KryptovaluuttaaJäljellä.Float64)
		}
		laskettuOstohinta = "NULL"
		if transaction.LaskettuOstohinta.Valid {
			laskettuOstohinta = fmt.Sprintf("'%f'", transaction.LaskettuOstohinta.Float64)
		}
		voitto = "NULL"
		if transaction.Voitto.Valid {
			voitto = fmt.Sprintf("'%f'", transaction.Voitto.Float64)
		}
		partialQuery = fmt.Sprintf(" ('%d', '%s', '%s', '%s', '%f', '%f', %s, %s, %s, '%s', '%f'),", total, transaction.Kryptovaluutta, transaction.Aikaleima.Format(time.RFC3339), transaction.Tyyppi, transaction.HintaEUR, transaction.MääräKryptovaluuttana, jaljella_krypto, laskettuOstohinta, voitto, transaction.Kommentti, transaction.EURPerKryptovaluutta)
		bulkQuery += partialQuery
	}
	bulkQuery = bulkQuery[:len(bulkQuery)-1] + ";"
	_, err = db.Exec(bulkQuery)
	if err != nil {
		return err
	}
	return nil
}
