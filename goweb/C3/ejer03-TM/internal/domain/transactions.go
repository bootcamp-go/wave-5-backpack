package domain

type Transaction struct {
	Id          int     `json:"id" binding:"required"`
	TranCode    string  `json:"tranCode" binding:"required"`
	Currency    string  `json:"currency" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Transmitter string  `json:"transmitter" binding:"required"`
	Reciever    string  `json:"reciever" binding:"required"`
	TranDate    string  `json:"tranDate" binding:"required"`
}
