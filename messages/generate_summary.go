package messages

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/stanlyzoolo/smartLaFamiliaBot/banks/commercial"
	"github.com/stanlyzoolo/smartLaFamiliaBot/banks/countries"
	"github.com/stanlyzoolo/smartLaFamiliaBot/banks/nacbank"
	"github.com/stanlyzoolo/smartLaFamiliaBot/services/myfin"
)

func GenerateSummaryForNatBank(rates []nacbank.Rate) (string, error) {
	header := fmt.Sprintf("Курс НацБанка РБ на %s.", time.Now().Format("02.01.2006"))

	report := make([]string, 0)

	for _, v := range rates {
		var b bytes.Buffer

		if v.Scale != 1 {
			t := template.Must(template.New("Russian rubles").Parse(RUB))
			if err := t.Execute(&b, v); err != nil {
				return "", fmt.Errorf("can't execute parsing data into template: %w", err)
			}

			report = append(report, b.String())
			b.Reset()
		} else {
			t := template.Must(template.New("USD and EUR").Parse(USDandEUR))
			if err := t.Execute(&b, v); err != nil {
				return "", fmt.Errorf("can't execute parsing data into template: %w", err)
			}

			report = append(report, b.String())
			b.Reset()
		}
	}

	return fmt.Sprint(header, strings.Join(report, "")), nil
}

func GenerateSummaryForCommercialBanks(rates []myfin.Currencies) (string, error) {
	header := fmt.Sprint("Курс коммерческих банков по данным Myfin.by\n")

	report := make([]string, len(rates))

	c := commercial.Rate{}
	var b bytes.Buffer

	for i, rate := range rates {
		c.Bank.Name = rate.BankName
		c.USD.Flag = countries.UnitedStates
		c.USD.Buying = rate.USDbuying
		c.USD.Selling = rate.USDselling
		c.EUR.Flag = countries.EuropeanUnion
		c.EUR.Buying = rate.EURbuying
		c.EUR.Selling = rate.EURselling
		c.RUB.Flag = countries.RussianFederation
		c.RUB.Buying = rate.RUBbuying
		c.RUB.Selling = rate.RUBselling
		c.ConvFromEURtoUSD.Buying = rate.EURtoUSDbuying
		c.ConvFromEURtoUSD.Selling = rate.EURtoUSDselling

		t := template.Must(template.New("Commercial").Parse(Commercial))
		if err := t.Execute(&b, c); err != nil {
			return "", fmt.Errorf("can't execute parsing data into template: %w", err)
		}

		report[i] = b.String()

		b.Reset()
	}

	return fmt.Sprint(header, strings.Join(report, "")), nil
}
