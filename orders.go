package gemini

import (
	"fmt"
	"net/http"
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


	requestURL := fmt.Sprintf("/v1/order/new")

	params := &Request{
			Url:		requestURL,
			nonce:		Nonce(),
			Symbol:		symbol,
			Amount:		fmt.Sprint(amount),
			Price:		fmt.Sprint(price),			
			Side:		side,
			OrderType:	"exchange limit",
		}		

  	_, err := c.Request("POST", requestURL, params, &order)

	return order, err
}

func (c *Client) CancelOrder(id int64) (*http.Response, error) {
	var order *http.Response


	requestURL := fmt.Sprintf("/v1/order/cancel")

	params := &Request{
			Url:		requestURL,
			nonce:		Nonce(),
			OrderID:	string(id),
		}		

  	_, err := c.Request("POST", requestURL, params, &order)

	return order, err
}

func (c *Client) CancelSessionOrders() (bool, error) {
	var result bool


	requestURL := fmt.Sprintf("/v1/order/cancel/session")

	params := &Request{
			Url:		requestURL,
			nonce:		Nonce(),
		}		

  	_, err := c.Request("POST", requestURL, params, &result)

	return result, err
}

func (c *Client) CancelAllOrders() (bool, error) {
	var result bool


	requestURL := fmt.Sprintf("/v1/order/cancel/all")

	params := &Request{
			Url:		requestURL,
			nonce:		Nonce(),
		}		

  	_, err := c.Request("POST", requestURL, params, &result)

	return result, err
}

func (c *Client) OrderStatus(id int64) (Order, error) {
	var order Order


	requestURL := fmt.Sprintf("/v1/order/status")

	params := &Request{
			Url:		requestURL,
			nonce:		Nonce(),
			OrderID:	string(id),
		}		

  	_, err := c.Request("POST", requestURL, params, &order)

	return order, err
}

func (c *Client) ActiveOrders() ([]Order, error) {
	var orders []Order


	requestURL := fmt.Sprintf("/v1/orders")

	params := &Request{
			Url:		requestURL,
			nonce:		Nonce(),
		}		

  	_, err := c.Request("POST", requestURL, params, &orders)

	return orders, err
}











