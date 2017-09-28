package gemini

import (
	"fmt"
)



type Order struct {
	Id		int64 		`json:"order_id"`
	ClientOrderId	int64		`json:"client_order_id"`
	Symbol		string		`json:"symbol"`
	Exchange	string		`json:"exchange"`
	price		float64		`json:",string"`
	AvgExPrice	float64		`json:"avg_execution_price,string"`
	Side		string		`json:"side"`
	Type		string		`json:"type"`
	options		[]string	`json:"options"`
	Timestamp	int64   	`json:"timestamp,string"`
	Timestamppms  	int64  		`json:"timestampms"`
	IsLive		bool		`json:"is_live"`
	IsCancelled	bool		`json:"is_cancelled"`
	WasForced	bool		`json:"was_forced"`
	ExecAmount	float64		`json:"executed_amount,string"`
	RemainAmount	float64		`json:"remaining_amount,string"`
	OrigAmount	float64		`json:"original_amount,string"`

}



func (c *Client) NewOrder(price float64, amount float64, side, symbol string) (Order, error) {
	var order Order


	requestURL := fmt.Sprintf("/order/new")

	params := &Request{
			Url:		requestURL,
			nonce:		Nonce(),
			Symbol:		symbol,
			Amount:		fmt.Sprint(amount),
			Price:		fmt.Sprint(price),			
			Side:		side,
			OrderType:	"exchange limit",
		}		

  	_, err := c.Request("GET", requestURL, params, &order)

	return order, err
}










