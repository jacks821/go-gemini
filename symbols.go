package gemini

import (
  "fmt"
)


func (c *Client) GetSymbols() ([]string, error) {
  var symbols []string
  requestURL := fmt.Sprintf("/symbols")

  _, err := c.Request("GET", requestURL, nil, &symbols)
  return symbols, err
}
