package transaction

import (
	"encoding/json"
	"fmt"
)

type Transaction struct {
	Id                                                           int
	TransactionCode, Currency, Emisor, Receiver, TransactionDate string
	Amount                                                       float64
}

func CreateTransactions() []Transaction {
	transactions := []Transaction{
		{Id: 1, TransactionCode: "1A", Currency: "COP", Emisor: "Santiago", Receiver: "Xartiago Inc.", TransactionDate: "06/07/2022", Amount: 30.50},
		{Id: 2, TransactionCode: "1B", Currency: "ARS", Emisor: "Rafael", Receiver: "Xarts S.A.S.", TransactionDate: "31/06/2022", Amount: 400.10},
		{Id: 3, TransactionCode: "1C", Currency: "CLP", Emisor: "Miguel", Receiver: "Tyago", TransactionDate: "20/03/2022", Amount: 92.13},
		{Id: 4, TransactionCode: "1D", Currency: "MXN", Emisor: "Angel", Receiver: "Pepito Store's", TransactionDate: "15/05/2022", Amount: 40.1},
	}
	return transactions
}

func Json() {
	transactions := CreateTransactions()
	transactionsJson, _ := json.Marshal(transactions)
	var reconverter []Transaction
	json.Unmarshal([]byte(transactionsJson), &reconverter)
	fmt.Println(reconverter)
}
