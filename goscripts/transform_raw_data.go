package main

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

type CurrencySet map[string]struct{}
type YearList []int

// transform the raw data from the database into the transaction struct
func TransformRawData(db *sql.DB) ([]Transaction, CurrencySet, YearList, error) {
	// Init variable
	var transaction Transaction
	var rawDataRow RawDataRow
	transactions := make([]Transaction, 0)
	currencySet := make(CurrencySet)
	yearList := make(YearList, 0)
	total := 0

	// Query database
	rows, err := db.Query("SELECT * FROM piggybank.crypto_csv_raw_data ORDER BY \"timestamp_UTC\" ASC")
	if err != nil {
		return []Transaction{}, nil, nil, err
	}
	defer rows.Close()

	// Loop through rows and fill transactions list
	for rows.Next() {
		total += 1
		//fmt.Println("Total rows:", total)
		rawDataRow, err = convertFromDBtoRawDataRow(rows)
		if err != nil {
			return []Transaction{}, nil, nil, err
		}

		transaction, err = convertToTransaction(rawDataRow)

		if err != nil {
			if err.Error() == "unknown transaction kind" {
				panic(err)
			}
			continue
		}
		currencySet[transaction.Kryptovaluutta] = struct{}{}
		transactions = append(transactions, transaction)
	}

	// Construct list of years
	var firstYear int = transactions[0].Aikaleima.Year()
	var lastYear int = rawDataRow.TimestampUTC.Year()
	for year := firstYear; year <= lastYear; year++ {
		yearList = append(yearList, year)
	}

	return transactions, currencySet, yearList, nil
}

func convertFromDBtoRawDataRow(rows *sql.Rows) (RawDataRow, error) {
	var rawDataRow RawDataRow
	err := rows.Scan(
		&rawDataRow.TimestampUTC,
		&rawDataRow.TransactionDesc,
		&rawDataRow.Currency,
		&rawDataRow.Amount,
		&rawDataRow.ToCurrency,
		&rawDataRow.ToAmount,
		&rawDataRow.NativeCurrency,
		&rawDataRow.NativeAmountUSD,
		&rawDataRow.TransactionKind,
		&rawDataRow.TransactionHash,
		&rawDataRow.NativeAmount,
		&rawDataRow.id,
	)
	return rawDataRow, err
}

func convertToTransaction(rawDataRow RawDataRow) (Transaction, error) {
	// Init
	var transaction Transaction

	// Condition
	transaction.Aikaleima = rawDataRow.TimestampUTC
	switch rawDataRow.TransactionKind {
	case "admin_wallet_credited":
		return Transaction{}, errors.New("skipping")
	case "card_cashback_reverted":
		return Transaction{}, errors.New("skipping")
	case "card_top_up":
		return Transaction{}, errors.New("skipping, topping up a card is not a taxable transaction")
	case "crypto_earn_interest_paid":
		transaction.Tyyppi = "DIVIDEND"
		transaction.Kryptovaluutta = rawDataRow.Currency
		transaction.HintaEUR = rawDataRow.NativeAmount
		transaction.MääräKryptovaluuttana = rawDataRow.Amount
		transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	case "crypto_earn_program_created":
		return Transaction{}, errors.New("skipping, creating an Earn deposit is not a taxable transaction")
	case "crypto_earn_program_withdrawn":
		return Transaction{}, errors.New("skipping, an Earn deposit ending is not a taxable transaction")
	case "crypto_purchase":
		transaction.Tyyppi = "BUY"
		transaction.Kryptovaluutta = rawDataRow.Currency
		transaction.HintaEUR = rawDataRow.NativeAmount
		transaction.MääräKryptovaluuttana = rawDataRow.Amount
		transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	case "crypto_viban_exchange":
		transaction.Tyyppi = "SELL"
		transaction.Kryptovaluutta = rawDataRow.Currency
		transaction.HintaEUR = rawDataRow.NativeAmount
		transaction.MääräKryptovaluuttana = -rawDataRow.Amount
		transaction.LaskettuOstohinta = sql.NullFloat64{Float64: 0.0, Valid: true}
	case "crypto_transfer":
		transaction.Tyyppi = "GIFT"
		transaction.Kryptovaluutta = rawDataRow.Currency
		transaction.HintaEUR = rawDataRow.NativeAmount
		transaction.MääräKryptovaluuttana = rawDataRow.Amount
		transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	case "crypto_wallet_swap_credited":
		return Transaction{}, errors.New("skipping")
	case "crypto_wallet_swap_debited":
		return Transaction{}, errors.New("skipping")
	case "finance.crypto_earn.loyalty_program_extra_interest_paid.crypto_wallet":
		transaction.Tyyppi = "DIVIDEND"
		transaction.Kryptovaluutta = rawDataRow.Currency
		transaction.HintaEUR = rawDataRow.NativeAmount
		transaction.MääräKryptovaluuttana = rawDataRow.Amount
		transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	case "finance.dpos.compound_interest.crypto_wallet":
		transaction.Tyyppi = "DIVIDEND"
		transaction.Kryptovaluutta = rawDataRow.Currency
		transaction.HintaEUR = rawDataRow.NativeAmount
		transaction.MääräKryptovaluuttana = rawDataRow.Amount
		transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	case "finance.dpos.staking.crypto_wallet":
		return Transaction{}, errors.New("skipping")
	case "finance.lockup.dpos_compound_interest.crypto_wallet":
		transaction.Tyyppi = "DIVIDEND"
		transaction.Kryptovaluutta = rawDataRow.Currency
		transaction.HintaEUR = rawDataRow.NativeAmount
		transaction.MääräKryptovaluuttana = rawDataRow.Amount
		transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	case "finance.lockup.dpos_lock.crypto_wallet":
		return Transaction{}, errors.New("skipping")
	case "lockup_lock":
		return Transaction{}, errors.New("skipping, locking up crypto is not a taxable transaction")
	case "lockup_unlock":
		return Transaction{}, errors.New("skipping, unlocking crypto is not a taxable transaction")
	case "lockup_upgrade":
		return Transaction{}, errors.New("skipping, upgrading locking is not a taxable transaction")
	case "mco_stake_reward":
		transaction.Tyyppi = "DIVIDEND"
		transaction.Kryptovaluutta = rawDataRow.Currency
		transaction.HintaEUR = rawDataRow.NativeAmount
		transaction.MääräKryptovaluuttana = rawDataRow.Amount
		transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	case "referral_bonus":
		return Transaction{}, errors.New("skipping")
	case "referral_card_cashback":
		transaction.Tyyppi = "CASHBACK"
		transaction.Kryptovaluutta = rawDataRow.Currency
		transaction.HintaEUR = rawDataRow.NativeAmount
		transaction.MääräKryptovaluuttana = rawDataRow.Amount
		transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	case "referral_gift":
		return Transaction{}, errors.New("skipping")
	case "reward.loyalty_program.trading_rebate.crypto_wallet":
		return Transaction{}, errors.New("skipping, volume too tiny, amount in USD only")
	case "rewards_platform_deposit_credited":
		transaction.Tyyppi = "DIVIDEND"
		transaction.Kryptovaluutta = rawDataRow.Currency
		transaction.HintaEUR = rawDataRow.NativeAmount
		transaction.MääräKryptovaluuttana = rawDataRow.Amount
		transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
	case "viban_purchase":
		transaction.Tyyppi = "BUY"
		transaction.Kryptovaluutta = rawDataRow.ToCurrency
		transaction.HintaEUR = rawDataRow.NativeAmount
		if !rawDataRow.ToAmount.Valid {
			return Transaction{}, errors.New("to amount is not valid")
		} else {
			transaction.MääräKryptovaluuttana = rawDataRow.ToAmount.Float64
			transaction.KryptovaluuttaaJäljellä = sql.NullFloat64{Float64: transaction.MääräKryptovaluuttana, Valid: true}
		}
		transaction.LaskettuOstohinta = sql.NullFloat64{Float64: 0.0, Valid: false}
	default:
		return Transaction{}, errors.New("unknown transaction kind")
	}
	transaction.EURPerKryptovaluutta = transaction.HintaEUR / transaction.MääräKryptovaluuttana
	return transaction, nil
}
