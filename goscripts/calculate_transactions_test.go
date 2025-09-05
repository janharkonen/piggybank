package main

import (
	"database/sql"
	"math"
	"testing"
)

func initTransactions() []Transaction {
	var transactions []Transaction
	var transaction Transaction

	// [0]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "BUY"
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: 267, Valid: true}
	transaction.HintaEUR = 100.04
	transaction.MääräKryptovaluuttana = 267
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [1]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "SELL"
	transaction.HintaEUR = 109.86
	transaction.MääräKryptovaluuttana = 267
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [2]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "BUY"
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: 300, Valid: true}
	transaction.HintaEUR = 105.51
	transaction.MääräKryptovaluuttana = 300
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [3]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "BUY"
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: 1000, Valid: true}
	transaction.HintaEUR = 198.47
	transaction.MääräKryptovaluuttana = 1000
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [4]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "SELL"
	transaction.HintaEUR = 463.87
	transaction.MääräKryptovaluuttana = 1300
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [5]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "BUY"
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: 1000, Valid: true}
	transaction.HintaEUR = 150.11
	transaction.MääräKryptovaluuttana = 1000
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [6]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "BUY"
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: 2000, Valid: true}
	transaction.HintaEUR = 400.22
	transaction.MääräKryptovaluuttana = 2000
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [7]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "SELL"
	transaction.HintaEUR = 51.01
	transaction.MääräKryptovaluuttana = 300
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [8]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "DOGE"
	transaction.Tyyppi = "SELL"
	transaction.HintaEUR = 256.78
	transaction.MääräKryptovaluuttana = 1500
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [9]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "CRO"
	transaction.Tyyppi = "BUY"
	transaction.HintaEUR = 1512.11
	transaction.MääräKryptovaluuttana = 12500
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [10]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "CRO"
	transaction.Tyyppi = "DIVIDEND"
	transaction.HintaEUR = 6.75
	transaction.MääräKryptovaluuttana = 58.74
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [11]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "CRO"
	transaction.Tyyppi = "DIVIDEND"
	transaction.HintaEUR = 7.75
	transaction.MääräKryptovaluuttana = 60
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [12]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "CRO"
	transaction.Tyyppi = "CASHBACK"
	transaction.HintaEUR = 0.12
	transaction.MääräKryptovaluuttana = 1.008
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [13]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "CRO"
	transaction.Tyyppi = "SELL"
	transaction.HintaEUR = 1320
	transaction.MääräKryptovaluuttana = 12000
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [14]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "CRO"
	transaction.Tyyppi = "SELL"
	transaction.HintaEUR = 57.86
	transaction.MääräKryptovaluuttana = 551.02
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [15]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "CRO"
	transaction.Tyyppi = "SELL"
	transaction.HintaEUR = 8.87
	transaction.MääräKryptovaluuttana = 68.221
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [16]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "CRO"
	transaction.Tyyppi = "BUY"
	transaction.HintaEUR = 217.5
	transaction.MääräKryptovaluuttana = 1500
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [17]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "CRO"
	transaction.Tyyppi = "GIFT"
	transaction.HintaEUR = 318
	transaction.MääräKryptovaluuttana = 2000
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [18]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "CRO"
	transaction.Tyyppi = "SELL"
	transaction.HintaEUR = 0.33
	transaction.MääräKryptovaluuttana = 3
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	transactions = append(transactions, transaction)

	// [19]
	transaction = Transaction{}
	transaction.Kryptovaluutta = "CRO"
	transaction.Tyyppi = "SELL"
	transaction.HintaEUR = 414.55
	transaction.MääräKryptovaluuttana = 2512.401
	transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
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

		// ----------------------------------------------------------------------------------
		sellIndex = 1
		calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows(transactions, "DOGE", sellIndex)
		if transactions[0].KryptovaluuttaaJäljellä.Float64 != 0 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if transactions[sellIndex].LaskettuOstohinta.Float64 != 100.04 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].Voitto.Float64-9.82) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].Verotettava.Float64-9.82) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}

		// ----------------------------------------------------------------------------------
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
		if math.Abs(transactions[sellIndex].Voitto.Float64-159.89) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].Verotettava.Float64-159.89) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}

		// ----------------------------------------------------------------------------------
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
		if math.Abs(transactions[sellIndex].Voitto.Float64-5.977) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].Verotettava.Float64-5.977) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}

		// ----------------------------------------------------------------------------------
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
		if math.Abs(transactions[sellIndex].Voitto.Float64+8.385) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].Verotettava.Float64+8.385) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}

		// ----------------------------------------------------------------------------------
		if transactions[9].KryptovaluuttaaJäljellä.Float64 != 12500 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		sellIndex = 13
		calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows(transactions, "CRO", sellIndex)
		if transactions[9].KryptovaluuttaaJäljellä.Float64 != 500 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].LaskettuOstohinta.Float64-1451.6256) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].Voitto.Float64+131.6256) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].Verotettava.Float64+131.6256) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}

		// ----------------------------------------------------------------------------------
		if transactions[10].KryptovaluuttaaJäljellä.Float64 != 58.74 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		sellIndex = 14
		calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows(transactions, "CRO", sellIndex)
		if transactions[9].KryptovaluuttaaJäljellä.Float64 != 0 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[10].KryptovaluuttaaJäljellä.Float64-7.72) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].LaskettuOstohinta.Float64-60.4844) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].Voitto.Float64+2.6244) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].Verotettava.Float64+2.6244) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}

		// ----------------------------------------------------------------------------------
		sellIndex = 15
		calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows(transactions, "CRO", sellIndex)
		if transactions[10].KryptovaluuttaaJäljellä.Float64 != 0 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if transactions[11].KryptovaluuttaaJäljellä.Float64 != 0 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[12].KryptovaluuttaaJäljellä.Float64-0.507) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].LaskettuOstohinta.Float64-0) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].Voitto.Float64-8.87) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].Verotettava.Float64-8.8103571429) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}

		// ----------------------------------------------------------------------------------
		if math.Abs(transactions[12].KryptovaluuttaaJäljellä.Float64-0.507) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		sellIndex = 18
		calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows(transactions, "CRO", sellIndex)
		if transactions[12].KryptovaluuttaaJäljellä.Float64 != 0 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if transactions[16].KryptovaluuttaaJäljellä.Float64 != 1497.507 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].LaskettuOstohinta.Float64-0.361485) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].Voitto.Float64+0.031485) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].Verotettava.Float64+0.0918421) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}

		// ----------------------------------------------------------------------------------
		if math.Abs(transactions[16].KryptovaluuttaaJäljellä.Float64-1497.507) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		sellIndex = 19
		calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows(transactions, "CRO", sellIndex)
		if transactions[16].KryptovaluuttaaJäljellä.Float64 != 0 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[17].KryptovaluuttaaJäljellä.Float64-985.106) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].LaskettuOstohinta.Float64-217.138515) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].Voitto.Float64-197.411485) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
		if math.Abs(transactions[sellIndex].Verotettava.Float64-36.043339) > 0.0000001 {
			t.Error("calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows test failed")
		}
	})
}
