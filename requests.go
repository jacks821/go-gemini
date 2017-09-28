package gemini

type Request struct {
	Url		string		`json:"request"`
	nonce		int64		`json:"nonce"`
	OrderID		string		`json:"client_order_id"`
	Symbol		string		`json:"symbol"`
	Amount		string		`json:"amount"`
	Price		string		`json:"price"`
	Side		string		`json:"side"`
	OrderType	string		`json:"type"`
	Options		[]string	`json:"options"`	

}
