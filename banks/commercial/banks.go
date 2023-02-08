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
		Flag            string
		Buying, Selling float32
	}

	EUR struct {
		Flag            string
		Buying, Selling float32
	}

	RUB struct {
		Flag            string
		Buying, Selling float32
	}

	ConvFromEURtoUSD struct {
		Buying, Selling float32
	}
)
