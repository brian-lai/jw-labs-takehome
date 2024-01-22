# Crypto Investor | JW Labs Takehome

## About
This is a CLI tool that prompts a user for $USD and spits out BTC and ETH to buy at the current rate.
It takes an amount in USD as input and calculates how much BTC and ETH to buy based on a 70/30 split.

### Update
There's a branch `web_app` that actually serves this sample functionality within a simple REST api.
I've chosen to use [mux](https://github.com/gorilla/mux) to handle routing
#### Usage
1. `go run cmd/main.go`
2. GET request to `http://localhost:8080/calculate-crypto-to-buy?amount=<amount-in-usd>`

The JSON response should yield the BTC and ETH to purchase

## Project Structure
- `crypto_investor.go`: The main entry point of the application. 
- `crypto_investor_test.go`:

## Coinbase API Wrapper
coinbase package is interface to interact with the Coinbase API.

- `client.go`: Implements the main API client, responsible for initializing and making HTTP requests to the Coinbase API.
- `models.go`: Contains data structures (models) for handling API responses. These models are designed to match the JSON response structure of the Coinbase API.
- `exchange_rates.go`: Handles specific API endpoint calls related to fetching exchange rates for cryptocurrencies.

## Setup
1. Clone the project repository to your local machine.
```
git clone git@github.com:brian-lai/jw-labs-takehome.git
```
2. Navigate to the project directory.
```
cd path/to/project
```

## Usage
To use the main application, run crypto_investor.go with the amount in USD you wish to invest:
```
go run crypto_investor.go
```

Enter amount in USD you want to invest when prompted.

## Running Tests
To run tests for the crypto_investor logic, execute the following command in the project's root directory:
```
go test
```

Use optional `-v` for verbose
