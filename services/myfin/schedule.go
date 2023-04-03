package myfin

func (s *service) run() {
	raw, err := s.scrapDomain()
	if err != nil {
		s.log.Error(err)
	}

	err = s.storeRates(raw)
	if err != nil {
		s.log.Error(err)
	}
}
