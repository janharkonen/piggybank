package main

import "fmt"

func MakeSummary(transactions []Transaction, currencySet CurrencySet, yearList YearList) {

	for currency := range currencySet {
		for _, year := range yearList {
			makeSummaryForCurrencyAndYear(transactions, currency, year)
		}
	}
}

func makeSummaryForCurrencyAndYear(transactions []Transaction, currency string, year int) {
	totalProfit := 0.0
	for _, transaction := range transactions {
		if transaction.Kryptovaluutta == currency && transaction.Aikaleima.Year() == year {
			if transaction.Tyyppi == "SELL" {
				totalProfit += transaction.Voitto.Float64
			}
		}
	}
	fmt.Println("Total profit for ", currency, "in", year, ":", totalProfit, "EUR")
}
