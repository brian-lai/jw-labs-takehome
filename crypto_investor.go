package main

import (
    "fmt"
    "strconv"
    "encoding/json"
    "jw-labs-takehome/coinbase"
    // "net/http"
    // "io/ioutil"
) // "jw-labs-takehome/coinbase is the path to the coinbase package in the GOPATH

func main() {
    // Check for command-line argument    
    fmt.Println("Enter amount in USD: ")

    var amountInput string

    fmt.Scan(&amountInput)

    // Parse the input amount
    amountInUSD, err := strconv.ParseFloat(amountInput, 64)

    if err != nil {
        fmt.Println("Error: Invalid amount: ", amountInput)
        return
    }

    fmt.Printf("Calculating amount of BTC and ETH to purchase for USD: $%v\n", amountInUSD)

    // Fetch exchange rates
    fmt.Println("Fetching exchange rates...")

    //////////////////////////////////////////////////////
    // var coinbaseAPI CoinbaseAPI
    // btcRate, ethRate, err := coinbaseAPI.FetchExchangeRates()
    // fmt.Printf("BTC Rate: %.6f (1 USD / 1 BTC)\n", btcRate)
    // fmt.Printf("ETH Rate: %.6f (1 USD / 1ETH)\n", ethRate)


    // Calculate amounts for BTC and ETH
    // btcAmount, ethAmount, err := coinbaseAPI.Calculate(amountInUSD, btcRate, ethRate)
    //////////////////////////////////////////////////////

    // using the coinbase package, the code above was replaced with the following:

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
}



/////////////////////////
// coinbase package will replace the code below
/////////////////////////

// // ExchangeRatesFetcher defines an interface for fetching exchange rates
// type ExchangeRatesFetcher interface {
//     FetchExchangeRates() (btcRate float64, ethRate float64, err error)
// }

// // CoinbaseAPI represents the Coinbase API 
// type CoinbaseAPI struct {

// }

// // FetchExchangeRates fetches exchange rates from Coinbase API
// func (c CoinbaseAPI) FetchExchangeRates() (btcRate float64, ethRate float64, err error) {
//     apiUrl := "https://api.coinbase.com/v2/exchange-rates?currency=USD"

//     resp, err := http.Get(apiUrl)
//     if err != nil {
//         return 0, 0, err
//     }
//     defer resp.Body.Close()

//     body, err := ioutil.ReadAll(resp.Body)
//     if err != nil {
//         return 0, 0, err
//     }

//     var rates struct {
//         Data struct {
//             Rates map[string]string `json:"rates"`
//         } `json:"data"`
//     }

//     if err := json.Unmarshal(body, &rates); err != nil {
//         return 0, 0, err
//     }

//     btcRate, err = strconv.ParseFloat(rates.Data.Rates["BTC"], 64)
//     if err != nil {
//         return 0, 0, err
//     }

//     ethRate, err = strconv.ParseFloat(rates.Data.Rates["ETH"], 64)
//     if err != nil {
//         return 0, 0, err
//     }

//     return btcRate, ethRate, nil
// }

// // Calculate returns the amount of BTC and ETH for a given amount in USD
// func (c CoinbaseAPI) Calculate(amountInUSD float64, btcRate float64, ethRate float64) (btcAmount float64, ethAmount float64, err error) {
//     if err != nil {
//         return 0, 0, err
//     }

//     btcAmount = (0.7 * amountInUSD) * btcRate
//     ethAmount = (0.3 * amountInUSD) * ethRate

//     return btcAmount, ethAmount, nil
// }
