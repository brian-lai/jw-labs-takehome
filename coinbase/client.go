package coinbase

import (
    "net/http"
    "fmt"
)

type Client struct {
    httpClient *http.Client
    apiKey     string
    baseURL    string
}

func NewClient(apiKey string) *Client {
    return &Client{
        httpClient: &http.Client{},
        apiKey:     apiKey,
        baseURL:    "https://api.coinbase.com/v2/",
    }
}

// Ping the Coinbase API (only used for testing really)
func (c *Client) Ping() error {
    // use the /time endpoint to ping the API
    resp, err := c.httpClient.Get(c.baseURL + "time")

    // Network error or unable to reach the base URL
    if err != nil {
        return err    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("Coinbase API is not reachable, status code: %d", resp.StatusCode)
    }

    return nil
}
