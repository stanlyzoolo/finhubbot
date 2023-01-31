package commercial

type (
	Rate struct {
		Bank             *Bank
		USD              *USD
		EUR              *EUR
		RUB              *RUB
		ConvFromEURtoUSD *ConvFromEURtoUSD
	}
	Bank struct {
		Name string
	}

	USD struct {
		Buying, Selling string
	}

	EUR struct {
		Buying, Selling string
	}

	RUB struct {
		Buying, Selling string
	}

	ConvFromEURtoUSD struct {
		Value string
	}
)
