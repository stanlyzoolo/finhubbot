package currencies

import (
	"fmt"
	"io"
	"net/http"
)

// func getURLenv() string {
// 	return os.Getenv("NBRB_CurrenciesRates")
// }

func CurrencyReq(url string, currencyID int) (*http.Request, error) {
	var body io.Reader

	urlWithCurrency := fmt.Sprint(url, currencyID)

	fmt.Println(urlWithCurrency)

	req, err := http.NewRequest(http.MethodGet, urlWithCurrency, body)
	if err != nil {
		return nil, fmt.Errorf("can't create request: %w", err)
	}
	// https://www.nbrb.by/api/exrates/rates/{431}

	// req.URL.Host = url

	// // TODO Можно использовать клиент: https://github.com/go-telegram-bot-api/telegram-bot-api/blob/v5.5.1/bot.go#L19
	// resp, err := tgbotapi.HTTPClient.Do(req)
	// if err != nil {
	// 	return nil, err
	// }

	// defer resp.Body.Close()

	// respBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }

	// var cur Currency
	// if err := json.Unmarshal(respBody, &cur); err != nil {
	// 	return nil, err
	// }

	return req, nil
}
