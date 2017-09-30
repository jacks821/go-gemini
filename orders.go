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
	IsLive		bool		`json:"is_live,string"`
	IsCancelled	bool		`json:"is_cancelled,string"`
	WasForced	bool		`json:"was_forced,string"`
	ExecAmount	float64		`json:"executed_amount,string"`
	RemainAmount	float64		`json:"remaining_amount,string"`
	OrigAmount	float64		`json:"original_amount,string"`

}



func (c *Client) NewOrder(price float64, amount float64, side, symbol string) (Order, error) {
	var order Order

	getNonce := Nonce()
	requestURL := fmt.Sprintf("/v1/order/new")

	params := &NewOrderRequest{
			Url:		requestURL,
			Nonce:		getNonce,
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

	getNonce := Nonce()
	requestURL := fmt.Sprintf("/v1/order/cancel")

	params := &OrderStatusRequest{
			Url:		requestURL,
			Nonce:		getNonce,
			OrderID:	string(id),
		}		

  	_, err := c.Request("POST", requestURL, params, &order)

	return order, err
}

func (c *Client) CancelSessionOrders() (bool, error) {
	var result bool
	
	getNonce := Nonce()

	requestURL := fmt.Sprintf("/v1/order/cancel/session")

	params := &BasicRequest{
			Url:		requestURL,
			Nonce:		getNonce,
		}		

  	_, err := c.Request("POST", requestURL, params, &result)

	return result, err
}

func (c *Client) CancelAllOrders() (bool, error) {
	var result bool
	
	getNonce := Nonce()

	requestURL := fmt.Sprintf("/v1/order/cancel/all")

	params := &BasicRequest{
			Url:		requestURL,
			Nonce:		getNonce,
		}		

  	_, err := c.Request("POST", requestURL, params, &result)

	return result, err
}

func (c *Client) OrderStatus(id int64) (Order, error) {
	var order Order
	
	getNonce := Nonce()

	requestURL := fmt.Sprintf("/v1/order/status")

	params := &OrderStatusRequest{
			Url:		requestURL,
			Nonce:		getNonce,
			OrderID:	string(id),
		}		

  	_, err := c.Request("POST", requestURL, params, &order)

	return order, err
}

func (c *Client) ActiveOrders() ([]Order, error) {
	var orders []Order

	getNonce := Nonce()

	requestURL := fmt.Sprintf("/v1/orders")

	params := &BasicRequest{
			Url:		requestURL,
			Nonce:		getNonce,
		}		

  	_, err := c.Request("POST", requestURL, params, &orders)

	return orders, err
}











