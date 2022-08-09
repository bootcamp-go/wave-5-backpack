package domain

type Transactions struct {
	Id              int64   `json:"id"`
	TransactionCode string  `json:"transaction_code"`
	TypeCurrency    string  `json:"type_of_currency"`
	Amount          float64 `json:"amount"`
	Transmitter     string  `json:"transmitter"`
	Receiver        string  `json:"receiver"`
	Date            string  `json:"date"`
	Completed       bool    `json:"completed"`
}
