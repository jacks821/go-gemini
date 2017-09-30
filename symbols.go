package gemini

import (
  "fmt"
)

//GetSymbols returns a list of the symbols actively traded on the Gemini exchange.
func (c *Client) GetSymbols() ([]string, error) {
  var symbols []string
  requestURL := fmt.Sprintf("/v1/symbols")

  _, err := c.Request("GET", requestURL, nil, &symbols)
  return symbols, err
}
