package main

import (
	"database/sql"
	"fmt"
)

// transform the raw data from the database into the transaction struct
func CalculateTransactions(transactions []Transaction, currencySet CurrencySet, yearList YearList) ([]Transaction, error) {

	calculateTransactionsForCurrency(transactions, "DOGE")
	//for currency := range currencySet {
	//	calculateTransactionsForCurrency(transactions, currency)
	//}
	return transactions, nil
}

func calculateTransactionsForCurrency(transactions []Transaction, currency string) {
	var sellIndexOffset int = -1
	for calculateNextSale(transactions, currency, &sellIndexOffset) {
		fmt.Println("id is:", sellIndexOffset+1)
	}
	fmt.Println(sellIndexOffset)
}

func calculateNextSale(
	transactions []Transaction,
	currency string,
	sellIndexOffset *int,
) bool {
	var sellIndexCandidate int = *sellIndexOffset + 1
	for sellIndexCandidate < len(transactions) {
		transaction := transactions[sellIndexCandidate]
		if transaction.Kryptovaluutta == currency && transaction.Tyyppi == "SELL" {
			*sellIndexOffset = sellIndexCandidate
			calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows(transactions, currency, sellIndexCandidate)
			return true
		}
		sellIndexCandidate++
	}
	*sellIndexOffset = -1
	return false

}

func calculateFIFOPurchasePriceAndAdjustCryptoBalanceInBuyRows(
	transactions []Transaction,
	currency string,
	sellIndex int,
) {
	var soldCryptoAmount float64 = transactions[sellIndex].MääräKryptovaluuttana
	var calucatedPurchasePrice float64 = 0
	var buyIndex int = 0
	for soldCryptoAmount > 0 {
		transaction := transactions[buyIndex]
		if (transaction.Kryptovaluutta == currency) && (transaction.Tyyppi == "BUY") && (transaction.KryptovaluuttaaJäljellä.Valid) && (transaction.KryptovaluuttaaJäljellä.Float64 != 0) {
			if soldCryptoAmount >= transaction.KryptovaluuttaaJäljellä.Float64 {
				soldCryptoAmount -= transaction.KryptovaluuttaaJäljellä.Float64
				calucatedPurchasePrice += transaction.HintaEUR
				transactions[buyIndex].KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: 0, Valid: true}
			}
		}
		buyIndex++
	}
	transactions[sellIndex].LaskettuOstohinta = sql.NullFloat64{Float64: calucatedPurchasePrice, Valid: true}
}
