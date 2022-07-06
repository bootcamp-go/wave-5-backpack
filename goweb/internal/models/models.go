package models

type Transaction struct {
  ID int `json:"id"`
  Cod string `json:"cod_transaction"`
  Moneda string `json:"moneda"`
  Monto float64 `json:"monto"`
  Emisor string `json:"emisor"`
  Receptor string `json:"receptor"`
  Fecha string `json:"date_transaction"`
}
