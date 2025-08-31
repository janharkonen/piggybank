package main

type currency struct {
	Name     string
	Symbol   string
	Decimals int
}

//out_indexes = [
//	'Kryptovaluutta',
//	'Aikaleima',
//	'Osto/Myynti',
//	'Hinta (EUR)',
//	'Määrä kryptovaluuttana',
//	'EUR/kryptovaluutta',
//	'Kryptovaluuttaa jäljellä (FIFO)',
//	'Laskettu ostohinta',
//	'Voitto',
//	'Kommentti'
//]

func main() {
	InitDbFromCsv("transactions.csv")
}
