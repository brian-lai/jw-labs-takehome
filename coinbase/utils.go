package coinbase

import (
    "strconv"
    "fmt"
)

// extract btcRate and ethRate from the response body
func ExtractRates(rates *ExchangeRates) (btcRate float64, ethRate float64, err error) {
    // check if BTC Rate exists
    _, ok := rates.Data.Rates["BTC"]
    if !ok {
        return 0, 0, fmt.Errorf("BTC Rate not found")
    }

    btcRate, err = strconv.ParseFloat(rates.Data.Rates["BTC"], 64)
    if err != nil {
        return 0, 0, err
    }

    // check if ETH Rate exists
    _, ok = rates.Data.Rates["ETH"]
    if !ok {
        return 0, 0, fmt.Errorf("ETH Rate not found")
    }

    ethRate, err = strconv.ParseFloat(rates.Data.Rates["ETH"], 64)
    if err != nil {
        return 0, 0, err
    }

    return btcRate, ethRate, nil
}

// Calculate the amount of BTC and ETH to buy given the 70/30 split
func Calculate(amountInUSD float64, btcRate float64, ethRate float64) (btcAmount float64, ethAmount float64, err error) {
    if err != nil {
        return 0, 0, err
    }

    btcAmount = (0.7 * amountInUSD) * btcRate
    ethAmount = (0.3 * amountInUSD) * ethRate

    return btcAmount, ethAmount, nil
}
