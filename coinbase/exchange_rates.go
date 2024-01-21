package coinbase

import (
    "encoding/json"
    // "fmt"
)

// use pointers so we mutate the same object (avoid copying)
func (c *Client) FetchExchangeRates() (*ExchangeRates, error) {
    resp, err := c.httpClient.Get(c.baseURL + "exchange-rates")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var rates ExchangeRates

    if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
        return nil, err
    }

    // fmt.Println(rates)

    return &rates, nil
}

