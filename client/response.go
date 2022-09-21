package client

// import (
// 	"encoding/json"
// 	"io"
// 	"net/http"

// 	"github.com/juju/errors"
// 	"github.com/stanlyzoolo/smartLaFamiliaBot/currencies"
// )

// func readResponse(resp *http.Response) (*currencies.Rate, error) {
// 	respBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, errors.Errorf("can't read response: %d", err)
// 	}

// 	defer resp.Body.Close()

// 	var cur currencies.Currency

// 	if err := json.Unmarshal(respBody, &cur); err != nil {
// 		return nil, errors.Errorf("can't unmarshal body: %d", err)
// 	}

// 	return &currencies.Rate{
// 		Currency: cur,
// 	}, nil
// }
