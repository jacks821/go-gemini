package gemini

import (
	"encoding/json"
)

type Request interface {
	GetRoute() string
	GetPayload()	[]byte
}


type NewOrderRequest struct {
	Url		string		`json:"request"`
	Nonce		int64		`json:"nonce,string"`
	Symbol		string		`json:"symbol"`
	Amount		string		`json:"amount"`
	Price		string		`json:"price"`
	Side		string		`json:"side"`
	OrderType	string		`json:"type"`
	Options		[]string	`json:"options"`	
}

type MyTradeRequest struct {
	Url		string		`json:"request"`
	Nonce		int64		`json:"nonce,string"`
	Symbol		string		`json:"symbol"`
	LimitTrades	int64		`json:"limit_trades,string"`
	Timestamp	int64 		`json:"timestamp,string"`
}


type BasicRequest struct {
	Url	string	`json:"request"`
	Nonce	int64	`json:"nonce,string"`
}

type OrderStatusRequest struct {
	Url	string	`json:"request"`
	Nonce	int64	`json:"nonce,string"`
	OrderID	string	`json:"order_id"`
}

func (r *OrderStatusRequest) GetRoute() string {
	return r.Url
}

func (r *BasicRequest) GetRoute() string {
	return r.Url
}

func (r *NewOrderRequest) GetRoute() string {
	return r.Url
}

func (r *OrderStatusRequest) GetPayload() []byte {
	data, _ := json.Marshal(r)
	return data
}

func (r *BasicRequest) GetPayload() []byte {
	data, _ := json.Marshal(r)
	return data
}

func (r *NewOrderRequest) GetPayload() []byte {
	data, _ := json.Marshal(r)
	return data
}

func (r *MyTradeRequest) GetPayload() []byte {
	data, _ := json.Marshal(r)
	return data
}

func (r *MyTradeRequest) GetRoute() string {
	return r.Url
}


