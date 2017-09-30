package gemini

import (
	"fmt"
	"net/http"
)


//Order will be the return of most of the orders endpoints.
type Order struct {
	ID		int64 		`json:"order_id"`
	ClientOrderID	int64		`json:"client_order_id"`
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


//NewOrder creates a new order on the Gemini exchange.
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

//CancelOrder cancels an existing Order that has not been filled. Requires the ID of the order to cancel.
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

//CancelSessionOrders cancels all orders created in a given session. Returns a boolean.
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

//CancelAllOrders cancels all outstanding Orders that have not been filled.
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

//OrderStatus returns the status of a given order by taking it's ID.
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

//ActiveOrders returns a list of the active orders.
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











