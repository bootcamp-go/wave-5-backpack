package requestdto

type TransactionRequest struct {
	CodTransaction string `json:"cod_transaction"`
	Currency        string `json:"currency"`
	Amount          int `json:"amount"`
	Sender          string `json:"sender"`
	Receiver        string `json:"receiver"`
	DateOrder      string `json:"date_order"`
}
