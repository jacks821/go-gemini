package gemini

import "fmt"

//TradeVolume includes all of the fields for a request for GetTradeVolume
type TradeVolume struct {
	AccountID		string		`json:"account_id"`
	Symbol			string		`json:"symbol"`
	BaseCurrency		float64		`json:"base_currency,string"`
	NotionalCurrency	float64		`json:"notional_currency,string"`
	DataDate		string		`json:"data_date"`
	TotalVolume		float64		`json:"total_volume_base,string"`
	MakerBuyRatio		float64		`json:"maker_buy_sell_ration,string"`
	BuyMakerBase		float64		`json:"buy_maker_base,string"`
	BuyMakerNotional	float64		`json:"buy_maker_notional,string"`
	BuyMakerCount		float64		`json:"buy_maker_count,string"`
	SellMakerBase		float64		`json:"sell_maker_base,string"`
	SellMakerCount		float64		`json:"sell_maker_count,string"`
	BuyTakerBase		float64		`json:"buy_taker_base,string"`
	BuyTakerNotional	float64		`json:"buy_taker_notional,string"`
	BuyTakerCount		float64		`json:"buy_taker_count,string"`
	SellTakerBase		float64		`json:"sell_taker_base,string"`
	SellTakerNotional	float64		`json:"sell_taker_notional,string"`
	SellTakerCount		float64		`json:"sell_taker_count,string"`
}

//Trade contains all the fields a Gemini trade should have
type Trade struct {
  Timestamp   int64 `json:"timestamp"`
  Timestamppms  int64  `json:"timestampms"`
  Tid       int64     `json:"tid"`
  Price     float64   `json:"price,string"`
  Amount    float64   `json:"amount,string"`
  Exchange    string  `json:"exchange"`
  Type        string    `json:"type"`
  Broken      bool      `json:"broken,string"`
}

//GetTrades gets a list of trades from Gemini.
func (c *Client) GetTrades(symbol string) ([]Trade, error) {
  var trades []Trade
  requestURL := fmt.Sprintf("/v1/trades/%s", symbol)

  _, err := c.Request("GET", requestURL, nil, &trades)
  return trades, err
}

//GetMyTrades returns your individual trades
func (c *Client) GetMyTrades(symbol string) ([]Trade, error) {
  	var trades []Trade
  	requestURL := fmt.Sprintf("/v1/mytrades")

	getNonce := Nonce()
	
	params := &MyTradeRequest {
		Url:		requestURL,
		Nonce:		getNonce,
		Symbol:		symbol,
		LimitTrades:	50,
	}	


  	_, err := c.Request("POST", requestURL, params, &trades)
  	return trades, err
}

//GetTradeVolume returns the trade volume on the exchange for up to 30 days.
func (c *Client) GetTradeVolume() ([]TradeVolume, error) {
	var volume []TradeVolume

	getNonce := Nonce()

	requestURL := fmt.Sprintf("/v1/tradevolume")

	params := &BasicRequest{
			Url:		requestURL,
			Nonce:		getNonce,
		}		

  	_, err := c.Request("POST", requestURL, params, &volume)

	return volume, err
}

