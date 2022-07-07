package domain

import "time"

type Transaction struct {
	Id              int       `json:"id" binding:"-"`
	TransactionCode string    `json:"transaction_code" binding:"-"`
	Currency        string    `json:"currency" binding:"required"`
	Amount          float64   `json:"amount" binding:"required"`
	Sender          string    `json:"sender" binding:"required"`
	Reciever        string    `json:"reciever" binding:"required"`
	TransactionDate time.Time `json:"transaction_date" binding:"-"`
}
