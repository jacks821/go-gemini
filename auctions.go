package gemini

import (
  "fmt"
)

type CurrentAuction struct {
  ClosedUntil   int64 `json:"closed_until_ms"`
  LastAuctionEID  int64  `json:"last_auction_eid"`
  LastAuctionPrice    float64  `json:",string"`
  LastAuctionQuantity   float64 `json:",string"`
  LastHighestBid        float64 `json:",string"`
  LastLowestAsk       float64 `json:",string"`
  MostRecentIndicativePrice   float64 `json:"most_recent_indicative_price"`
  MostRecentIndicativeQuantity  float64 `json:"most_recent_indicative_quantity"`
  MostRecentHighestBid      float64 `json:"most_recent_highest_bid_price"`
  MostRecentLowestAsk       float64 `json:"most_recent_lowest_ask_price"`
  NextUpdate          int64 `json:"next_update_ms"`
  NextAuction         int64 `json:"next_auction_ms"`
}

type HistoricalAuction struct {
  Timestamp   int64 `json:"timestamp"`
  Timestamppms  int64  `json:"timestampms"`
  AuctionID     int64   `json:"auction_id"`
  EID         int64     `json:"eid"`
  EventType   string    `json:"event_type"`
  Result      string    `json:"auction_result"`
  Price       float64   `json:",string"`
  Quantity     float64  `json:",string"`
  HighestBid   float64  `json:",string"`
  LowestAsk    float64  `json:",string"`
}


func (c *Client) GetCurrentAuction(symbol string) (CurrentAuction, error) {
  var auction CurrentAuction
  requestURL := fmt.Sprintf("/v1/auction/%s", symbol)

  _, err := c.Request("GET", requestURL, nil, &auction)
  return auction, err
}

func (c *Client) GetHistoricalAuction(symbol string) ([]HistoricalAuction, error) {
  var auctions []HistoricalAuction
  requestURL := fmt.Sprintf("/v1/auction/%s/history", symbol)

  _, err := c.Request("GET", requestURL, nil, &auctions)
  return auctions, err
}
