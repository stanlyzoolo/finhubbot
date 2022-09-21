package client

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/juju/errors"
)

// Формат запроса в НБ РБ:
// https://www.nbrb.by/api/exrates/rates/840?parammode=1
func getRequest(url string, curID int) (*http.Request, error) {
	var body io.Reader

	ctx := context.Background()

	urlWithCurrency := fmt.Sprint(url, curID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlWithCurrency, body)
	if err != nil {
		return nil, errors.Errorf("can't set request: %d", err)
	}

	q := req.URL.Query()
	q.Add("parammode", "1")

	req.URL.RawQuery = q.Encode()

	return req, nil
}
