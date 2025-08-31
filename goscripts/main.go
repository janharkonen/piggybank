package main

import "fmt"

type Transaction struct {
	Kryptovaluutta          string
	Aikaleima               string
	Tyyppi                  string // "BUY", "SELL", "CASHBACK", "DIVIDEND"
	HintaEUR                float64
	MääräKryptovaluuttana   float64
	EURKryptovaluutta       float64
	KryptovaluuttaaJäljellä float64
	LaskettuOstohinta       float64
	Voitto                  float64
	Kommentti               string
}

func main() {
	db, err := GetDbClient()
	if err != nil {
		fmt.Println("Error getting database:", err)
	}
	defer db.Close()

	err = InitDbFromCsv(db, "transactions.csv")
	if err != nil {
		fmt.Println("Error initializing database from CSV:", err)
	}
	//transactions, err := TransformRawData()
	//if err != nil {
	//	fmt.Println("Error transforming raw data:", err)
	//	return
	//}
	//fmt.Println(transactions)
}
