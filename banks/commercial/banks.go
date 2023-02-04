package commercial

type (
	Rate struct {
		Bank             Bank
		USD              USD
		EUR              EUR
		RUB              RUB
		ConvFromEURtoUSD ConvFromEURtoUSD
	}
	Bank struct {
		Name string
	}

	USD struct {
		Flag, Buying, Selling string
	}

	EUR struct {
		Flag, Buying, Selling string
	}

	RUB struct {
		Flag, Buying, Selling string
	}

	ConvFromEURtoUSD struct {
		Buying, Selling string
	}
)
