# Crypto Investor Project

## Introduction
The Crypto Investor Project is a Go application designed to simplify investing in cryptocurrencies like Bitcoin (BTC) and Ethereum (ETH). It includes a command-line tool (`crypto_investor.go`) for calculating investment splits and a Coinbase API wrapper (`coinbase` package) for fetching real-time cryptocurrency exchange rates.

## Project Structure
- `crypto_investor.go`: The main entry point of the application. It takes an amount in USD as input and calculates how much BTC and ETH to buy based on a 70/30 split.
- `crypto_investor_test.go`:

## Coinbase API Wrapper
coinbase package is interface to interact with the Coinbase API.

- `client.go`: Implements the main API client, responsible for initializing and making HTTP requests to the Coinbase API.
- `models.go`: Contains data structures (models) for handling API responses. These models are designed to match the JSON response structure of the Coinbase API.
- `exchange_rates.go`: Handles specific API endpoint calls related to fetching exchange rates for cryptocurrencies.

## Setup
1. Clone the project repository to your local machine.
```
git clone [repository-url]
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
