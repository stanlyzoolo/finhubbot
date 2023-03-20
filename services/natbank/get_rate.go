package natbank

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

func (c *service) getCurrenciesRates(ctx context.Context) ([]Rate, error) {
	// TODO потому что валют всего 3
	rates := make([]Rate, 0)

	for code, flag := range codesAndFlags {
		req, err := preparetRequest(ctx, c.cfg.NBRB.OneRateURL, code)
		if err != nil {
			c.log.Error(err)

			return nil, fmt.Errorf("%w", err)
		}

		resp, err := c.client.Do(req)
		if err != nil {
			c.log.Errorf("can't Do request: %v", err)

			return nil, fmt.Errorf("can't Do request: %w", err)
		}

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			c.log.Errorf("can't read response: %v", err)

			return nil, fmt.Errorf("can't read response: %w", err)
		}

		defer resp.Body.Close()

		var rate Rate

		if err := json.Unmarshal(respBody, &rate); err != nil {
			return nil, fmt.Errorf("can't unmarshal body: %w", err)
		}

		rate.Icon = flag

		rates = append(rates, rate)
	}

	return rates, nil
}
