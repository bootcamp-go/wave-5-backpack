package domain

type Transanction struct {
	Id       int     `json:"id"`
	Code     string  `json:"code"`
	Coin     string  `json:"coin"`
	Amount   float64 `json:"amount"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
	Date     string  `json:"date"`
}
