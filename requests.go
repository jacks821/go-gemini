package gemini

import (
	"encoding/json"
)

//Request is an interface which makes it simpler to pass Requests to the Request function.
type Request interface {
	GetRoute() string
	GetPayload()	[]byte
}

//NewOrderRequest is a struct which contains all the fields required to create a new order.
type NewOrderRequest struct {
	URL		string		`json:"request"`
	Nonce		int64		`json:"nonce,string"`
	Symbol		string		`json:"symbol"`
	Amount		string		`json:"amount"`
	Price		string		`json:"price"`
	Side		string		`json:"side"`
	OrderType	string		`json:"type"`
	Options		[]string	`json:"options"`	
}

//MyTradeRequest is a struct which contains all the fields required to view your trades.
type MyTradeRequest struct {
	URL		string		`json:"request"`
	Nonce		int64		`json:"nonce,string"`
	Symbol		string		`json:"symbol"`
	LimitTrades	int64		`json:"limit_trades,string"`
	Timestamp	int64 		`json:"timestamp,string"`
}

//BasicRequest is the smallest amount of information which one can use to generate a request to the Gemini API.  Is used in a few methodsd.
type BasicRequest struct {
	URL	string	`json:"request"`
	Nonce	int64	`json:"nonce,string"`
}

//OrderStatusRequest contains all the fields needed to 
type OrderStatusRequest struct {
	URL	string	`json:"request"`
	Nonce	int64	`json:"nonce,string"`
	OrderID	string	`json:"order_id"`
}

//GetRoute returns the URL of the request.  Required to implement the Request interface.
func (r *OrderStatusRequest) GetRoute() string {
	return r.URL
}

//GetRoute returns the URL of the request.  Required to implement the Request interface.
func (r *BasicRequest) GetRoute() string {
	return r.URL
}

//GetRoute returns the URL of the request.  Required to implement the Request interface.
func (r *NewOrderRequest) GetRoute() string {
	return r.URL
}

//GetPayload returns the data payload.  Required to implement the Request interface.
func (r *OrderStatusRequest) GetPayload() []byte {
	data, _ := json.Marshal(r)
	return data
}

//GetPayload returns the data payload.  Required to implement the Request interface.
func (r *BasicRequest) GetPayload() []byte {
	data, _ := json.Marshal(r)
	return data
}

//GetPayload returns the data payload.  Required to implement the Request interface.
func (r *NewOrderRequest) GetPayload() []byte {
	data, _ := json.Marshal(r)
	return data
}

//GetPayload returns the data payload.  Required to implement the Request interface.
func (r *MyTradeRequest) GetPayload() []byte {
	data, _ := json.Marshal(r)
	return data
}

//GetRoute returns the URL of the request.  Required to implement the Request interface.
func (r *MyTradeRequest) GetRoute() string {
	return r.URL
}


