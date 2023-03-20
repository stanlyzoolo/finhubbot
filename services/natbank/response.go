package natbank

type Rate struct {
	Currency
	Icon UTFCountryIcon
}

type Currency struct {
	ID           int     `json:"Cur_ID"`           // внутренний код НЦ РБ
	Abbreviation string  `json:"Cur_Abbreviation"` // буквенный код: USD
	Name         string  `json:"Cur_Name"`         // наименование валюты на русском языке во множественном,
	Scale        int     `json:"Cur_Scale"`        // количество единиц валюты
	OfficialRate float32 `json:"Cur_OfficialRate"` // курс
	// Date         time.Time `json:"Date"`             // TODO будет активно после кастомного анмаршалера - https://habr.com/ru/post/492996/
}
