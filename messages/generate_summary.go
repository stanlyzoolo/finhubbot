package messages

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/juju/errors"

	"github.com/stanlyzoolo/smartLaFamiliaBot/banks/nacbank"
)

func GenerateFromNatBankRates(rates []nacbank.Rate) (string, error) {
	header := fmt.Sprintf("Курс Национального Банка РБ на %s.", time.Now().Format("02.01.2006"))

	report := make([]string, 0)

	for _, v := range rates {
		var b bytes.Buffer

		if v.Scale != 1 {
			t := template.Must(template.New("MessageTemplateScales").Parse(MessageTemplateScales))
			if err := t.Execute(&b, v); err != nil {
				return "", errors.Errorf("can't execute parsing data into template: %d", err)
			}

			report = append(report, b.String())
			b.Reset()
		} else {
			t := template.Must(template.New("MessageTemplateScale").Parse(MessageTemplateScale))
			if err := t.Execute(&b, v); err != nil {
				return "", errors.Errorf("can't execute parsing data into template: %d", err)
			}

			report = append(report, b.String())
			b.Reset()
		}
	}

	return fmt.Sprint(header, strings.Join(report, "")), nil
}
