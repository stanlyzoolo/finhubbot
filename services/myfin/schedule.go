package myfin

import "time"

// TODO можно создать отдельный тип Schedule и заполнить его интервалами
func (s *service) runBySchedule() {
	weekdays := s.allowedWeekdays()
	delay := time.Second * time.Duration(s.cfg.MyFin.Delay)
	every := time.Second * time.Duration(s.cfg.MyFin.Every)
	iterations := s.cfg.MyFin.IterationsCount

	ticker := time.NewTicker(time.Second)

	s.run(ticker, delay, every, iterations, weekdays)
}

func (s *service) run(t *time.Ticker, delay, every time.Duration, iterations int, allowedWeekdays *Weekdays) {
	done := make(chan bool)

	defer t.Stop()

	for {
		select {
		case <-done:
			s.log.Info("Done")

			return
		case <-t.C:
			time.Sleep(delay)

			today := time.Now().Weekday().String()

			if allowedWeekdays.Has(today) {
				for i := 1; i <= iterations; i++ {
					// Myfin
					raw, err := s.scrapDomain()
					if err != nil {
						s.log.Error(err)
					}

					err = s.storeRates(raw)
					if err != nil {
						s.log.Error(err)
					}

					time.Sleep(every)
				}
			}
		}
	}
}

func (s *service) allowedWeekdays() *Weekdays {
	weekdays := make(Weekdays, len(s.cfg.MyFin.AllowedWeekdays))

	for i, day := range s.cfg.MyFin.AllowedWeekdays {
		weekdays[day] = Weekday(i)
	}

	return &weekdays
}
