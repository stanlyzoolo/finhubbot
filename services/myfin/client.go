package myfin

import (
	"fmt"

	"github.com/stanlyzoolo/smartLaFamiliaBot/config"
	"github.com/stanlyzoolo/smartLaFamiliaBot/log"

	"github.com/gocolly/colly"
)

type Myfin interface {
	SetAllowedDomain(domain string)
	ScrapDomain() []string
	OrderIncomingData(in []string) (ordered []BankCurrencies)
}

type service struct {
	c   *colly.Collector
	log *log.Logger
	cfg *config.Config
}

type BankCurrencies struct {
	Bank            string
	USDbuying       string
	USDselling      string
	EURbuying       string
	EURselling      string
	RUBbuying       string
	RUBselling      string
	EURtoUSDbuying  string
	EURtoUSDselling string
}

func New(log *log.Logger, cfg *config.Config) Myfin {
	return &service{
		c:   colly.NewCollector(),
		log: log,
		cfg: cfg,
	}
}

func (s *service) SetAllowedDomain(domain string) {
	s.log.Infof("Set domain: %s", domain)

	s.c.AllowedDomains = []string{domain}
}

func (s *service) ScrapDomain() (banks []string) {
	// TODO возможно так не будет работать
	// banks = make([]string, 0)

	// Все банки
	s.c.OnHTML(`tbody[class="sort_body"]`, func(h *colly.HTMLElement) {
		h.ForEach(`tr[class="c-currency-table__main-row c-currency-table__main-row--with-arrow"]`, func(i int, h *colly.HTMLElement) {
			h.ForEach("td", func(i int, h *colly.HTMLElement) {
				banks = append(banks, h.Text)
			})
		})
	})

	return
}

func (s *service) OrderIncomingData(in []string) (ordered []BankCurrencies) {
	ordered = orderBanksDetails(in)

	return
}

func orderBanksDetails(raw []string) []BankCurrencies {
	var (
		bank  BankCurrencies
		banks []BankCurrencies
	)

	for {
		if len(raw) < 9 {
			break
		}

		bank.Bank = raw[0]
		bank.USDbuying = raw[1]
		bank.USDselling = raw[2]
		bank.EURbuying = raw[3]
		bank.EURselling = raw[4]
		bank.RUBbuying = raw[5]
		bank.RUBselling = raw[6]
		bank.EURtoUSDbuying = raw[7]
		bank.EURtoUSDselling = raw[8]

		fmt.Printf("Ready bank info: %+v\n", bank)

		raw = raw[9:]

		fmt.Println("Next iteration")

		banks = append(banks, bank)
	}

	return banks
}

// func sh() {
// 	bank := make([]string, 0)

// 	c := colly.NewCollector(colly.AllowedDomains("myfin.by"))

// 	// bankDetails := c.Clone()
// 	// Все банки
// 	c.OnHTML(`tbody[class="sort_body"]`, func(h *colly.HTMLElement) {
// 		h.ForEach(`tr[class="c-currency-table__main-row c-currency-table__main-row--with-arrow"]`, func(i int, h *colly.HTMLElement) {
// 			h.ForEach("td", func(i int, h *colly.HTMLElement) {
// 				bank = append(bank, h.Text)
// 			})
// 		})
// 	})

// 	c.OnRequest(func(r *colly.Request) {
// 		log.Println("Visiting: ", r.URL.String())
// 	})

// 	if err := c.Visit("https://myfin.by/currency/minsk"); err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(orderBanksDetails(bank))
// }
