package myfin

import "github.com/gocolly/colly"

func (s *service) scrapDomain() ([]string, error) {
	banks := make([]string, 0)

	// Все коммерческие банки
	s.c.OnHTML(`tbody[class="sort_body"]`, func(h *colly.HTMLElement) {
		h.ForEach(`tr[class="c-currency-table__main-row c-currency-table__main-row--with-arrow"]`, func(i int, h *colly.HTMLElement) {
			h.ForEach("td", func(i int, h *colly.HTMLElement) {
				banks = append(banks, h.Text)
			})
		})
	})

	// s.c.OnRequest(func(r *colly.Request) {
	// 	s.log.Infof("visiting: %s", r.URL.String())
	// })

	// флаг повторного посещения ресурса
	s.c.AllowURLRevisit = true

	err := s.c.Visit(s.cfg.MyFin.URL)
	if err != nil {
		s.log.Errorf("can't visit filled url: %v", err)

		return nil, err
	}

	return banks, err
}
