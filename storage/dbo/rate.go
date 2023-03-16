package dbo

import "database/sql"

type (
	Rate struct {
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
)
