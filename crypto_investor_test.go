package main

import (
    "testing"
    "jw-labs-takehome/coinbase"
    "strconv"
    "encoding/json"
)

// Use a mock exchange rate provider to test the ExtractRates() function
type MockExchangeRates struct {
    BTCRate float64
    ETHRate float64
}

func (m MockExchangeRates) GetExchangeRates() (*coinbase.ExchangeRates, error) {
    rates := &coinbase.ExchangeRates{
        Data: struct {
            Rates map[string]string `json:"rates"`
        }{
            Rates: map[string]string{
                "BTC": strconv.FormatFloat(m.BTCRate, 'f', -1, 64),
                "ETH": strconv.FormatFloat(m.ETHRate, 'f', -1, 64),
            },
        },
    }
    return rates, nil
}

// test the extract function
func TestExtractRates(t *testing.T) {
    // use mock exchange rates
    exchangeRatesMocker := MockExchangeRates{
        BTCRate: 0.000024, 
        ETHRate: 0.000405,
    }

    exchangeRates, err := exchangeRatesMocker.GetExchangeRates()
    if err != nil {
        // this shouldn't fail since we're using a mock
        t.Fatalf("GetExchangeRates() error = %v", err)
    }

    // extract the rates from the mock structure (should mimic the API response)
    btcRate, ethRate, err := coinbase.ExtractRates(exchangeRates)
    if err != nil {
        t.Fatalf("ExtractRates() error = %v", err)
    }

    if btcRate != 0.000024 || ethRate != 0.000405 {
        t.Errorf("ExtractRates() = %v, %v; want %v, %v", btcRate, ethRate, 0.000024, 0.000405)
    }
}

// test the calculate function
func TestCalculate(t *testing.T) {
    amountInUSD := 1000.0
    btcRate := 0.00002 // Example rate
    ethRate := 0.00030 // Example rate

    expectedBTC := (0.7 * amountInUSD) * btcRate
    expectedETH := (0.3 * amountInUSD) * ethRate

    btcAmount, ethAmount, err := coinbase.Calculate(amountInUSD, btcRate, ethRate)
    if err != nil {
        t.Fatalf("Calculate() error = %v", err)
    }

    if btcAmount != expectedBTC || ethAmount != expectedETH {
        t.Errorf("Calculate() = %v, %v; want %v, %v", btcAmount, ethAmount, expectedBTC, expectedETH)
    }
}

// test the api client
func TestClient(t *testing.T) {
    // random API key since this API doesn't require authentication
    apiKey := "random-api-key"

    // instantiate new client
    client := coinbase.NewClient(apiKey)

    // send a ping to the API and see if it responds
    err := client.Ping()
    if err != nil {
        t.Fatalf("Ping() error = %v", err)
    }
}

// test the exchange_rates endpoint
func TestExchangeRates(t *testing.T) {
    // create new client
    client := coinbase.NewClient("random-api-key")

    // get exchange rate
    exchangeRates, err := client.FetchExchangeRates()

    if err != nil {
        t.Fatalf("GetExchangeRates() error = %v", err)
    }

    if exchangeRates.Data.Rates["BTC"] == "" || exchangeRates.Data.Rates["ETH"] == "" {
        t.Errorf("GetExchangeRates() = %v, %v; want %v, %v", exchangeRates.Data.Rates["BTC"], exchangeRates.Data.Rates["ETH"], "non-empty", "non-empty")
    }
}

// test the models
func TestExchangeRatesUnmarshal(t *testing.T) {
    // example JSON response from the API
    jsonStr := `{"data": {"currency": "USD", "rates": {"BTC": "0.00002", "ETH": "0.00030"}}}`

    // unmarshal the JSON into ExchangeRates struct
    var rates coinbase.ExchangeRates
    err := json.Unmarshal([]byte(jsonStr), &rates)
    if err != nil {
        t.Fatalf("Unmarshal failed: %v", err)
    }

    // Assert fields in rates...
    if rates.Data.Rates["BTC"] != "0.00002" || rates.Data.Rates["ETH"] != "0.00030" {
        t.Errorf("Unmarshal() = %v, %v; want %v, %v", rates.Data.Rates["BTC"], rates.Data.Rates["ETH"], "0.00002", "0.00030")
    }
}
