package nacbank

type CodeByISO4217 int
type DigitModeByISO4217 int

// Коды валют по ИСО 4217
const (
	USD CodeByISO4217 = 840
	EUR CodeByISO4217 = 978
	RUB CodeByISO4217 = 643

	digitMode DigitModeByISO4217 = 1
)

type UTFCountryIcon string

const (
	UnitedStates      UTFCountryIcon = "\U0001F1FA\U0001F1F8"
	EuropeanUnion     UTFCountryIcon = "\U0001F1EA\U0001F1FA"
	RussianFederation UTFCountryIcon = "\U0001F1F7\U0001F1FA"
	ChinaRepublic     UTFCountryIcon = "\U0001F1E8\U0001F1F3"
)

var CodesAndFlags = map[CodeByISO4217]UTFCountryIcon{
	USD: UnitedStates,
	EUR: EuropeanUnion,
	RUB: RussianFederation,
}
