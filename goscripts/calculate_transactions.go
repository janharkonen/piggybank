package main

import (
	"database/sql"
	"fmt"
)

// transform the raw data from the database into the transaction struct
func CalculateTransactions(transactions []Transaction, currencySet CurrencySet, yearList YearList) ([]Transaction, error) {

	for currency := range currencySet {
		calculateTransactionsForCurrency(transactions, currency)
	}
	return transactions, nil
}

func calculateTransactionsForCurrency(transactions []Transaction, currency string) {
	var sellIndexOffset int = -1
	for calculateNextSale(transactions, currency, &sellIndexOffset) {
		fmt.Println("SELL", currency, ":", sellIndexOffset+1)
	}
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
	if transactions[sellIndex].Tyyppi != "SELL" {
		panic("sell index is not a sell transaction")
	}
	var soldCryptoAmount float64 = transactions[sellIndex].MääräKryptovaluuttana
	var calculatedPurchasePrice float64 = 0
	var buyIndex int = 0
	for soldCryptoAmount > 0 {
		transaction := transactions[buyIndex]

		if isAnAqcuiredAsset(transaction, currency) {
			if soldCryptoAmount >= transaction.KryptovaluuttaaJäljellä.Float64 {
				calculatedPurchasePrice += transaction.KryptovaluuttaaJäljellä.Float64 * transaction.EURPerKryptovaluutta
				transactions[buyIndex].KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: 0, Valid: true}
				soldCryptoAmount -= transaction.KryptovaluuttaaJäljellä.Float64
			} else {
				calculatedPurchasePrice += soldCryptoAmount * transaction.EURPerKryptovaluutta
				transactions[buyIndex].KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.KryptovaluuttaaJäljellä.Float64 - soldCryptoAmount, Valid: true}
				soldCryptoAmount = 0
			}
		}
		buyIndex++
	}
	transactions[sellIndex].LaskettuOstohinta = sql.NullFloat64{Float64: calculatedPurchasePrice, Valid: true}
	profit := transactions[sellIndex].HintaEUR - calculatedPurchasePrice
	transactions[sellIndex].Voitto = sql.NullFloat64{Float64: profit, Valid: true}
}

func isAnAqcuiredAsset(transaction Transaction, currency string) bool {

	isRightCurrency := transaction.Kryptovaluutta == currency
	if !isRightCurrency {
		return false
	}

	isAnAcquiredAsset := transaction.Tyyppi == "BUY"
	if !isAnAcquiredAsset {
		return false
	}

	transactionHasCryptoLeftToSell := transaction.KryptovaluuttaaJäljellä.Valid && transaction.KryptovaluuttaaJäljellä.Float64 != 0
	return transactionHasCryptoLeftToSell
}
