package main

import (
    log "log"
    http "net/http"
    mux "github.com/gorilla/mux"
    handlers "jw-labs-takehome/pkg/handlers"
) // "jw-labs-takehome/coinbase is the path to the coinbase package in the GOPATH

func main() {
    router := mux.NewRouter()

    // GET route to fetch how much crypto to buy
    // should pull USD from query string
    router.HandleFunc("/calculate-crypto-to-buy", handlers.CalculateCryptoToBuy).Methods("GET")

    log.Println("API server listening on port 8080")
    http.ListenAndServe(":8080", router)
}
