package gemini

import (
  "fmt"
)

//Orders is a set of bids and asks
type Orders struct{
  Bids []Bid `json:",string"`
  Asks []Ask  `json:",string"`
}

//Bid includes the price and amount of the bid
type Bid struct {
  Price float64   `json:",string"`
  Amount float64  `json:",string"`
  Timestamp string `json:"timestamp"`
}

//Ask includes the price and amount of the ask
type Ask struct {
  Price float64   `json:",string"`
  Amount float64  `json:",string"`
  Timestamp string `json:"timestamp"`
}

//GetBook gets a list of the current orders on the market as an array of bids and asks.
func (c *Client) GetBook(symbol string) (Orders, error) {
  var orders Orders
  requestURL := fmt.Sprintf("/v1/book/%s", symbol)

  _, err := c.Request("GET", requestURL, nil, &orders)
  return orders, err
}
