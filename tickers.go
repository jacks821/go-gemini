package gemini

import (
  "fmt"
)

type Ticker struct {
  Bid     float64 `json:",string"`
  Ask     float64 `json:",string"`
  Last    float64 `json:",string"`
  Volume  Volume  `json:",string"`

}

type Volume struct {
  Timestamp       int64   `json:"timestamp"`
  BTC             float64     `json:",string"`
  ETH             float64     `json:",string"`
  USD             float64     `json:",string"`
}

func (c *Client) GetTicker(symbol string) (Ticker, error) {
  var ticker Ticker
  requestURL := fmt.Sprintf("/v1/pubticker/%s", symbol)

  _, err := c.Request("GET", requestURL, nil, &ticker)
  return ticker, err
}
