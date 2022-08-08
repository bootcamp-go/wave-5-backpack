package models

type Transaction struct {
	ID       int     `json:"id"`
	Monto    float64 `json:"monto"`
	Cod      string  `json:"cod_transaction"`
	Moneda   string  `json:"moneda"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
	Fecha    string  `json:"date_transaction"`
}
