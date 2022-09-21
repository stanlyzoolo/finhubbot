package currencies

type Rate struct {
	Currency
	CountryFlagUTF8 string
}

// Флаги в UTF-8
const (
	UnitedStates      = "\U0001F1FA\U0001F1F8"
	EuropeanUnion     = "\U0001F1EA\U0001F1FA"
	RussianFederation = "\U0001F1F7\U0001F1FA"
	ChinaRepublic     = "\U0001F1E8\U0001F1F3"
)
