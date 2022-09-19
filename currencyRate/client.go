package currencies

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func ShowCurrency(url string) (*Currency, error) {
// 	var (
// 		body io.Reader
// 		// ctx  context.Context
// 	)

// 	req, err := http.NewRequest(http.MethodGet, url, body)
// 	if err != nil {
// 		return nil, fmt.Errorf("Can't create request: %w", err)
// 	}

// 	// TODO Можно использовать клиент: https://github.com/go-telegram-bot-api/telegram-bot-api/blob/v5.5.1/bot.go#L19
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer resp.Body.Close()

// 	respBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var cur Currency
// 	if err := json.Unmarshal(respBody, &cur); err != nil {
// 		return nil, err
// 	}

// 	return &cur, nil
// }
