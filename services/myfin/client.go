package myfin

import (
	"fmt"

	"github.com/stanlyzoolo/smartLaFamiliaBot/config"
	"github.com/stanlyzoolo/smartLaFamiliaBot/log"

	"github.com/gocolly/colly"
)

type Myfin interface {
	SetAllowedDomain() error
	ScrapDomain() ([]string, error)
	OrderIncomingData(in []string) (ordered []Currencies)
}

type service struct {
	c   *colly.Collector
	log *log.Logger
	cfg *config.Config
}

type Currencies struct {
	BankName        string
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

func (s *service) SetAllowedDomain() error {
	domain := s.cfg.MyFin.AllowedDomain
	url := s.cfg.MyFin.URL

	if domain == "" || url == "" {
		return fmt.Errorf("allowed domain and url are not set in env's")
	}

	s.c.AllowedDomains = []string{domain}
	s.cfg.MyFin.URL = url

	s.log.Infof("current config for MyFin: %s, %s", domain, url)

	return nil
}

func (s *service) ScrapDomain() ([]string, error) {
	banks := make([]string, 0)

	// Все коммерческие банки
	s.c.OnHTML(`tbody[class="sort_body"]`, func(h *colly.HTMLElement) {
		h.ForEach(`tr[class="c-currency-table__main-row c-currency-table__main-row--with-arrow"]`, func(i int, h *colly.HTMLElement) {
			h.ForEach("td", func(i int, h *colly.HTMLElement) {
				banks = append(banks, h.Text)
			})
		})
	})

	s.c.OnRequest(func(r *colly.Request) {
		s.log.Infof("Visiting: %s", r.URL.String())
	})

	// флаг повторного посещения ресурса
	s.c.AllowURLRevisit = true

	err := s.c.Visit(s.cfg.MyFin.URL)
	if err != nil {
		s.log.Errorf("can't visit filled url: %v", err)

		return nil, err
	}

	return banks, nil
}

func (s *service) OrderIncomingData(in []string) (ordered []Currencies) {
	ordered = orderBanksDetails(in)

	return
}

func orderBanksDetails(raw []string) []Currencies {
	var (
		bank  Currencies
		banks []Currencies
	)

	for {
		if len(raw) < 9 {
			break
		}

		bank.BankName = raw[0]
		bank.USDbuying = raw[1]
		bank.USDselling = raw[2]
		bank.EURbuying = raw[3]
		bank.EURselling = raw[4]
		bank.RUBbuying = raw[5]
		bank.RUBselling = raw[6]
		bank.EURtoUSDbuying = raw[7]
		bank.EURtoUSDselling = raw[8]

		raw = raw[9:]

		banks = append(banks, bank)
	}

	return banks
}
