package gemini

import (
  "fmt"
)

type Orders struct{
  Bids []Bid `json:",string"`
  Asks []Ask  `json:",string"`
}

type Bid struct {
  Price float64   `json:",string"`
  Amount float64  `json:",string"`
  Timestamp string `json:"timestamp"`
}

type Ask struct {
  Price float64   `json:",string"`
  Amount float64  `json:",string"`
  Timestamp string `json:"timestamp"`
}

func (c *Client) GetBook(symbol string) (Orders, error) {
  var orders Orders
  requestURL := fmt.Sprintf("/book/%s", symbol)

  _, err := c.Request("GET", requestURL, nil, &orders)
  return orders, err
}
