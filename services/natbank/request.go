package natbank

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// Формат запроса в НБ РБ:
// https://www.nbrb.by/api/exrates/rates/840?parammode=1
func preparetRequest(ctx context.Context, url string, code CodeByISO4217) (*http.Request, error) {
	var body io.Reader

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprint(url, code), body)
	if err != nil {
		return nil, fmt.Errorf("can't prepare request: %w", err)
	}

	q := req.URL.Query()
	q.Add("parammode", fmt.Sprint(digitMode))

	req.URL.RawQuery = q.Encode()

	return req, err
}
