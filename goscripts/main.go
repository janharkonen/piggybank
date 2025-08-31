package main

import (
	"database/sql"
	"fmt"
	"time"
)

type Transaction struct {
	Kryptovaluutta          string
	Aikaleima               time.Time
	Tyyppi                  string // "BUY", "SELL", "CASHBACK", "DIVIDEND"
	HintaEUR                float64
	MääräKryptovaluuttana   float64
	EURPerKryptovaluutta    float64
	KryptovaluuttaaJäljellä float64
	LaskettuOstohinta       float64
	Voitto                  float64
	Kommentti               string
}

type RawDataRow struct {
	TimestampUTC    time.Time
	TransactionDesc string
	Currency        string
	Amount          float64
	ToCurrency      string
	ToAmount        sql.NullFloat64
	NativeCurrency  string
	NativeAmount    float64
	NativeAmountUSD float64
	TransactionKind string
	TransactionHash string
	id              int
}

func main() {
	db, err := GetDbClient()
	if err != nil {
		fmt.Println("Error getting database:", err)
	}
	defer db.Close()

	//err = InitDbFromCsv(db, "transactions.csv")
	//if err != nil {
	//	fmt.Println("Error initializing database from CSV:", err)
	//}

	transactions, currencySet, yearList, err := TransformRawData(db)
	if err != nil {
		fmt.Println("Error transforming raw data:", err)
		return
	}
	fmt.Println("--------------------------------")
	fmt.Println(currencySet)
	fmt.Println(yearList)

	fmt.Println(transactions[0])
}
