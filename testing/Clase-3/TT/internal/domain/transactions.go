package domain

type Transaction struct {
	Id          int     `json:"id"`
	TranCode    string  `json:"tranCode"`
	Currency    string  `json:"currency"`
	Amount      float64 `json:"amount"`
	Transmitter string  `json:"transmitter"`
	Receiver    string  `json:"receiver"`
	TranDate    string  `json:"tranDate"`
}
