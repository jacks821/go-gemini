package gemini

import (
  "fmt"
)

//Ticker is an item which is returned from the GetTicker method. It includes bids and asks.
type Ticker struct {
  Bid     float64 `json:",string"`
  Ask     float64 `json:",string"`
  Last    float64 `json:",string"`
  Volume  Volume  `json:",string"`

}


//Volume is a struct which is returned in the Ticker item. It shows the volumes in each currency.
type Volume struct {
  Timestamp       int64   `json:"timestamp"`
  BTC             float64     `json:",string"`
  ETH             float64     `json:",string"`
  USD             float64     `json:",string"`
}


//GetTicker returns a ticker for a symbol which includes the highest bid and lowest ask currently available.
func (c *Client) GetTicker(symbol string) (Ticker, error) {
  var ticker Ticker
  requestURL := fmt.Sprintf("/v1/pubticker/%s", symbol)

  _, err := c.Request("GET", requestURL, nil, &ticker)
  return ticker, err
}
