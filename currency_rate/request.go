package currencyrate

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/juju/errors"
)

func CurrencyReq(url string, currencyID int) (*http.Request, error) {
	var body io.Reader

	ctx := context.Background()

	urlWithCurrency := fmt.Sprint(url, currencyID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlWithCurrency, body)
	if err != nil {
		return nil, errors.Errorf("can't set request: %d", err)
	}

	return req, nil
}
