package main

import (
	"database/sql"
	"fmt"
	"math"
)

// transform the raw data from the database into the transaction struct
func CalculateTransactions(transactions []Transaction, currencySet CurrencySet, yearList YearList) ([]Transaction, error) {

	for currency := range currencySet {
		calculateTransactionsForCurrency(transactions, currency)
	}
	return transactions, nil
}

func calculateTransactionsForCurrency(transactions []Transaction, currency string) {
	fmt.Println("Calculating transactions for currency:", currency, "start")
	var sellIndexOffset int = -1
	for calculateNextSale(transactions, currency, &sellIndexOffset) {
		continue
		//fmt.Println("SELL", currency, ":", sellIndexOffset+1)
	}
	fmt.Println("Calculating transactions for currency:", currency, "done")
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
	var totalCalculatedPurchasePrice float64 = 0
	var calculatedPurchasePrice float64 = 0
	var buyIndex int = 0
	for soldCryptoAmount > 0 {
		if buyIndex >= len(transactions) {
			panic("no more resources to even sell")
		}
		transaction := transactions[buyIndex]
		if isAnAqcuiredAsset(transaction, currency) {

			if (transaction.Tyyppi != "CASHBACK") && (transaction.Tyyppi != "DIVIDEND") && (transaction.Tyyppi != "GIFT") {
				calculatedPurchasePrice = math.Min(transaction.KryptovaluuttaaJäljellä.Float64, soldCryptoAmount) * transaction.EURPerKryptovaluutta
			} else {
				calculatedPurchasePrice = 0
			}
			if soldCryptoAmount >= transaction.KryptovaluuttaaJäljellä.Float64 {
				totalCalculatedPurchasePrice += calculatedPurchasePrice
				transactions[buyIndex].KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: 0, Valid: true}
				soldCryptoAmount -= transaction.KryptovaluuttaaJäljellä.Float64
			} else {
				totalCalculatedPurchasePrice += calculatedPurchasePrice
				transactions[buyIndex].KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.KryptovaluuttaaJäljellä.Float64 - soldCryptoAmount, Valid: true}
				soldCryptoAmount = 0
			}
		}
		buyIndex++
	}
	transactions[sellIndex].LaskettuOstohinta = sql.NullFloat64{Float64: totalCalculatedPurchasePrice, Valid: true}
	profit := transactions[sellIndex].HintaEUR - totalCalculatedPurchasePrice
	transactions[sellIndex].Voitto = sql.NullFloat64{Float64: profit, Valid: true}
}

func isAnAqcuiredAsset(transaction Transaction, currency string) bool {

	isRightCurrency := transaction.Kryptovaluutta == currency
	if !isRightCurrency {
		return false
	}

	isAnAcquiredAsset := (transaction.Tyyppi == "BUY") || (transaction.Tyyppi == "CASHBACK") || (transaction.Tyyppi == "DIVIDEND") || (transaction.Tyyppi == "GIFT")
	if !isAnAcquiredAsset {
		return false
	}

	transactionHasCryptoLeftToSell := transaction.KryptovaluuttaaJäljellä.Valid && transaction.KryptovaluuttaaJäljellä.Float64 != 0
	return transactionHasCryptoLeftToSell
}
