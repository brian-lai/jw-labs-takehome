package main

import (
    "testing"
)

// MockExchangeRatesFetcher is a mock implementation for testing
type MockExchangeRatesFetcher struct{}

func (m MockExchangeRatesFetcher) FetchExchangeRates() (float64, float64, error) {
    return 50000.0, 4000.0, nil // Mocked rates for BTC and ETH
}

func TestCalculateCryptoAmounts(t *testing.T) {
    btcRate, ethRate := .00002401, .00040487
    amountInUSD := 1000.0
    expectedBTC := (0.7 * amountInUSD) * btcRate
    expectedETH := (0.3 * amountInUSD) * ethRate

    // CoinbaseAPI implements the ExchangeRatesFetcher interface
    coinbaseAPI := CoinbaseAPI{}

    btcAmount, ethAmount, err := coinbaseAPI.Calculate(amountInUSD, btcRate, ethRate)

    if btcAmount != expectedBTC || ethAmount != expectedETH || err != nil {
        t.Errorf("CoinbaseAPI.Calculate() = %v, %v; want %v, %v", btcAmount, ethAmount, expectedBTC, expectedETH)
    }
}

// Test ExchangeRatesFetcher interface
func TestFetchExchangeRates(t *testing.T) {
    mockExchangeRatesFetcher := MockExchangeRatesFetcher{}

    btcRate, ethRate, err := mockExchangeRatesFetcher.FetchExchangeRates()

    if err != nil {
        t.Errorf("FetchExchangeRates() returned an error: %v", err)
    }

    if btcRate != 50000.0 || ethRate != 4000.0 {
        t.Errorf("FetchExchangeRates() = %v, %v; want %v, %v", btcRate, ethRate, 50000.0, 4000.0)
    }
}


// Test CoinbaseAPI implementation
func TestCoinbaseAPI(t *testing.T) {
    coinbaseAPI := CoinbaseAPI{}

    btcRate, ethRate, err := coinbaseAPI.FetchExchangeRates()

    if err != nil {
        t.Errorf("FetchExchangeRates() returned an error: %v", err)
    }

    if btcRate == 0.0 || ethRate == 0.0 {
        t.Errorf("FetchExchangeRates() = %v, %v; want non-zero values", btcRate, ethRate)
    }
}
