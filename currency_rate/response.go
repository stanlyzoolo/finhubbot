package currencyrate

// import "time"

type Currency struct {
	ID int64 `json:"Cur_ID"` // внутренний код НЦ РБ
	Abbreviation string `json:"Cur_Abbreviation"` // буквенный код: USD
	Scale        int64  `json:"Cur_Scale"`        // количество единиц валюты
	Name         string `json:"Cur_Name"`         // наименование валюты на русском языке во множественном,
	// либо в единственном числе, в зависимости от количества единиц
	OfficialRate float32 `json:"Cur_OfficialRate"` // курс
	Flag         string
	// Date         time.Time `json:"Date"` // Don`t need
}
