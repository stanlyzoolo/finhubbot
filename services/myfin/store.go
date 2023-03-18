package myfin

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/stanlyzoolo/smartLaFamiliaBot/storage/dbo"
)

func (s *service) storeRates(raw []string) error {
	currencies := orderBanksDetails(raw)

	models := s.mapToRateModel(currencies)

	var err error

	for _, model := range models {
		err = s.storage.Create(context.Background(), model)
		if err != nil {
			return err
		}
	}

	return err
}

func orderBanksDetails(raw []string) []Currency {
	var (
		bank  Currency
		banks []Currency
	)

	for {
		// название банка + 6 значений куров USD, EUR, RUB
		if len(raw) < 7 {
			break
		}

		bank.Bank.Name = raw[0]
		bank.USD.Buying = raw[1]
		bank.USD.Selling = raw[2]
		bank.EUR.Buying = raw[3]
		bank.EUR.Selling = raw[4]
		bank.RUB.Buying = raw[5]
		bank.RUB.Selling = raw[6]

		raw = raw[7:]

		banks = append(banks, bank)
	}

	return banks
}

func (s *service) mapToRateModel(currencies []Currency) []dbo.CommercialRate {
	dboRates := make([]dbo.CommercialRate, len(currencies))

	dboRate := dbo.CommercialRate{}
	for i, c := range currencies {
		dboRate.Bank.Name = sql.NullString{
			String: c.Bank.Name,
			Valid:  len(c.Bank.Name) != 0,
		}
		dboRate.USD.Buying = toFloat32(c.USD.Buying)
		dboRate.USD.Selling = toFloat32(c.USD.Selling)
		dboRate.EUR.Buying = toFloat32(c.EUR.Buying)
		dboRate.EUR.Selling = toFloat32(c.EUR.Selling)
		dboRate.RUB.Buying = toFloat32(c.RUB.Buying)
		dboRate.RUB.Selling = toFloat32(c.RUB.Selling)

		dboRates[i] = dboRate
	}

	return dboRates
}

func toFloat32(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)

	if err != nil {
		return 0
	}

	return float32(f)
}
