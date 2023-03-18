package dbo

import "database/sql"

type (
	CommercialRate struct {
		Bank Bank
		USD  USD
		EUR  EUR
		RUB  RUB
	}
	Bank struct {
		Name sql.NullString
	}

	USD struct {
		Buying, Selling float32
	}

	EUR struct {
		Buying, Selling float32
	}

	RUB struct {
		Buying, Selling float32
	}

	NatBankRate struct {
		ID           int
		Abbreviation sql.NullString
		Name         sql.NullString
		Scale        int
		OfficialRate float32
	}
)
