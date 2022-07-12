package domain

type Transaction struct {
	Id        int     `json:"id"`
	Code      string  `json:"code" binding:"required"`
	Currency  string  `json:"currency" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
	Issuer    string  `json:"issuer" binding:"required"`
	Recipient string  `json:"recipient" binding:"required"`
	Date      string  `json:"date" binding:"required"`
}
