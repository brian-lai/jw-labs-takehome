package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "jw-labs-takehome/coinbase"
)

func CalculateCryptoToBuy(w http.ResponseWriter, r *http.Request) {

    // Check for command-line argument    
    fmt.Println("Enter amount in USD: ")

    // Read input amount from http request
    var amountInput string = r.URL.Query().Get("amount")

    // fmt.Scan(&amountInput)

    // Parse the input amount
    amountInUSD, err := strconv.ParseFloat(amountInput, 64)

    if err != nil {
        fmt.Println("Error: Invalid amount: ", amountInput)
        return
    }

    fmt.Printf("Calculating amount of BTC and ETH to purchase for USD: $%v\n", amountInUSD)

    // Fetch exchange rates
    fmt.Println("Fetching exchange rates...")

    // instantiate a new Coinbase API client
    coinbaseAPIClient := coinbase.NewClient("api-key-that-does-not-matter")

    // fetch exchange rates
    rates, err := coinbaseAPIClient.FetchExchangeRates()
    fmt.Println("Exchange rates fetched.")

    // extract btcRate and ethRate from the response body
    btcRate, ethRate, err := coinbase.ExtractRates(rates)

    // var coinbaseAPI CoinbaseAPI
    // btcRate, ethRate, err := coinbaseAPI.FetchExchangeRates()
    fmt.Printf("BTC Rate: %.6f (1 USD / 1 BTC)\n", btcRate)
    fmt.Printf("ETH Rate: %.6f (1 USD / 1ETH)\n", ethRate)


    // Calculate amounts for BTC and ETH
    btcAmount, ethAmount, err := coinbase.Calculate(amountInUSD, btcRate, ethRate)

    // Create a JSON response
    var response = map[string]float64{
        "BTC": btcAmount,
        "ETH": ethAmount,
    }

    // Print JSON response to STDOUT
    json, err := json.Marshal(response)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(json))

    // Write JSON response to http response
    w.Header().Set("Content-Type", "application/json")
    w.Write(json)
}
