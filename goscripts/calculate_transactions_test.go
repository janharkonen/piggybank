package main

import (
	"database/sql"
	"math"
	"testing"
)

func initTransactions() []Transaction {
	var transactions []Transaction
	var transaction Transaction

	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "BUY"
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: 267, Valid: true}
	transaction.HintaEUR = 100.04
	transaction.MääräKryptovaluuttana = 267
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "SELL"
	transaction.HintaEUR = 109.86
	transaction.MääräKryptovaluuttana = 267
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "BUY"
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: 300, Valid: true}
	transaction.HintaEUR = 105.51
	transaction.MääräKryptovaluuttana = 300
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "BUY"
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: 1000, Valid: true}
	transaction.HintaEUR = 198.47
	transaction.MääräKryptovaluuttana = 1000
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "SELL"
	transaction.HintaEUR = 463.87
	transaction.MääräKryptovaluuttana = 1300
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "BUY"
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: 1000, Valid: true}
	transaction.HintaEUR = 150.11
	transaction.MääräKryptovaluuttana = 1000
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "BUY"
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: 2000, Valid: true}
	transaction.HintaEUR = 400.22
	transaction.MääräKryptovaluuttana = 2000
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "SELL"
	transaction.HintaEUR = 51.01
	transaction.MääräKryptovaluuttana = 300
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "SELL"
	transaction.HintaEUR = 256.78
	transaction.MääräKryptovaluuttana = 1500
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	return transactions
}

func TestCalculateTransactions(t *testing.T) {

	t.Run("dummy test 1", func(t *testing.T) {
		if 1 != 1 {
			t.Error("Dummy test 1 failed")
		}
	})

	t.Run("Transaction datatype test", func(t *testing.T) {
		transactions := initTransactions()
		if transactions[0].Kryptovaluutta != "DOGE" {
			t.Error("Transaction datatype test failed")
		}
		if transactions[0].KryptovaluuttaaJäljellä.Float64 != 267 {
			t.Error("Transaction datatype test failed")
		}
	})
	t.Run("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test", func(t *testing.T) {
		transactions := initTransactions()
		var sellIndex int

		sellIndex = 1
		calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows(transactions, "DOGE", sellIndex)
		if transactions[0].KryptovaluuttaaJäljellä.Float64 != 0 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if transactions[sellIndex].LaskettuOstohinta.Float64 != 100.04 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}

		sellIndex = 4
		calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows(transactions, "DOGE", sellIndex)
		if transactions[2].KryptovaluuttaaJäljellä.Float64 != 0 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if transactions[4].KryptovaluuttaaJäljellä.Float64 != 0 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if transactions[sellIndex].LaskettuOstohinta.Float64 != 303.98 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}

		sellIndex = 7
		calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows(transactions, "DOGE", sellIndex)
		if transactions[5].KryptovaluuttaaJäljellä.Float64 != 700 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if transactions[6].KryptovaluuttaaJäljellä.Float64 != 2000 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].LaskettuOstohinta.Float64-45.033) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}

		sellIndex = 8
		calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows(transactions, "DOGE", sellIndex)
		if transactions[5].KryptovaluuttaaJäljellä.Float64 != 0 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if transactions[6].KryptovaluuttaaJäljellä.Float64 != 1200 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].LaskettuOstohinta.Float64-265.165) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
	})
}
