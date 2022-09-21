package messages

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/juju/errors"

	"github.com/stanlyzoolo/smartLaFamiliaBot/currencies"
)

func (s *Summary) GenerateFromRates(rates []currencies.Rate) (string, error) {
	s.date = time.Now().Format("02.01.2006")
	s.header = fmt.Sprintf("Курс Национального Банка РБ на %s.", s.date)

	for _, v := range rates {
		var b bytes.Buffer

		if v.Scale != 1 {
			t := template.Must(template.New("MessageTemplateScales").Parse(MessageTemplateScales))
			if err := t.Execute(&b, v); err != nil {
				return "", errors.Errorf("can't execute parsing data into template: %d", err)
			}

			s.report = append(s.report, b.String())
			b.Reset()
		} else {
			t := template.Must(template.New("MessageTemplateScale").Parse(MessageTemplateScale))
			if err := t.Execute(&b, v); err != nil {
				return "", errors.Errorf("can't execute parsing data into template: %d", err)
			}

			s.report = append(s.report, b.String())
			b.Reset()
		}
	}

	return fmt.Sprint(s.header, strings.Join(s.report, "")), nil
}
