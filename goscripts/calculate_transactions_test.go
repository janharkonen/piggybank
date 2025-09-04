package main

import (
	"database/sql"
	"testing"
)

func TestCalculateTransactions(t *testing.T) {
	t.Run("dummy test 1", func(t *testing.T) {
		if 1 != 1 {
			t.Error("Dummy test 1 failed")
		}
	})

	t.Run("Transaction datatype test", func(t *testing.T) {

		var transactions []Transaction
		var transaction Transaction

		transaction = Transaction{}
		transaction.Kryptovaluutta = "DOGE"
		transaction.Tyyppi = "BUY"
		transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: 267, Valid: true}
		transaction.HintaEUR = 100.04
		transaction.MääräKryptovaluuttana = 267
		transactions = append(transactions, transaction)

		transaction = Transaction{}
		transaction.Kryptovaluutta = "DOGE"
		transaction.Tyyppi = "SELL"
		transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: 267, Valid: true}
		transaction.HintaEUR = 100.04
		transaction.MääräKryptovaluuttana = 267
		transactions = append(transactions, transaction)

		if 2 != 2 {
			t.Error("Dummy test 2 failed")
		}
	})
}
