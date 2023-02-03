package nacbank

import (
	"github.com/stanlyzoolo/smartLaFamiliaBot/banks/countries"
)

type Rate struct {
	Currency
	CountryFlagUTF8 string
}

type Currency struct {
	ID           int64   `json:"Cur_ID"`           // внутренний код НЦ РБ
	Abbreviation string  `json:"Cur_Abbreviation"` // буквенный код: USD
	Name         string  `json:"Cur_Name"`         // наименование валюты на русском языке во множественном,
	Scale        int64   `json:"Cur_Scale"`        // количество единиц валюты
	OfficialRate float32 `json:"Cur_OfficialRate"` // курс
	// Date         time.Time `json:"Date"`             // TODO будет активно после кастомного анмаршалера - https://habr.com/ru/post/492996/
}

// Коды валют по ИСО 4217
const (
	USD = 840
	EUR = 978
	RUS = 643
)

var CodesAndFlags = map[int]string{
	USD: countries.UnitedStates,
	EUR: countries.EuropeanUnion,
	RUS: countries.RussianFederation,
}
