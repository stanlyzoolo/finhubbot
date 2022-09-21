package currencies

type Currency struct {
	ID           int64   `json:"Cur_ID"`           // внутренний код НЦ РБ
	Abbreviation string  `json:"Cur_Abbreviation"` // буквенный код: USD
	Name         string  `json:"Cur_Name"`         // наименование валюты на русском языке во множественном,
	Scale        int64   `json:"Cur_Scale"`        // количество единиц валюты
	OfficialRate float32 `json:"Cur_OfficialRate"` // курс
	// Date         time.Time `json:"Date"` // Don`t need
}

// Коды валют по ИСО 4217
const (
	USD = 840
	EUR = 978
	RUS = 643
	CNY = 156
)

var CodesAndFlags = map[int]string{
	USD: UnitedStates,
	EUR: EuropeanUnion,
	RUS: RussianFederation,
	CNY: ChinaRepublic,
}
