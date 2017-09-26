package gemini

import "fmt"

type Trade struct {
  Timestamp   int64 `json:"timestamp"`
  Timestamppms  int64  `json:"timestampms"`
  Tid       int64     `json:"tid"`
  Price     float64   `json:",string"`
  Amount    float64   `json:",string"`
  Exchange    string  `json:"exchange"`
  Type        string    `json:"type"`
  Broken      bool      `json:",string"`
}

func (c *Client) GetTrades(symbol string) ([]Trade, error) {
  var trades []Trade
  requestURL := fmt.Sprintf("/trades/%s", symbol)

  _, err := c.Request("GET", requestURL, nil, &trades)
  return trades, err
}
