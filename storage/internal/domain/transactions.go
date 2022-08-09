package domain

type Transaction struct {
	Id              int `json:"id"`
	CodTransaction string `json:"codTransaction"`
	Currency        string `json:"currency"`
	Amount          int `json:"amount"`
	Sender          string `json:"sender"`
	Receiver        string `json:"receiver"`
	DateOrder      string `json:"dateOrder"`
}
