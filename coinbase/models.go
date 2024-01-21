package coinbase

// Example of a model for exchange rate data
type ExchangeRates struct {
    // Data structure from the Coinbase API is 
    // {"data":{"currency":"USD","rates":{"BTC": "1.00","ETH": "1.00"}}}
    Data struct {
        Rates map[string]string `json:"rates"`
    } `json:"data"`
}

