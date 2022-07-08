package models

type Transaction struct {
  ID int `json:"id"`
  Cod string `json:"cod_transaction" binding:"required"`
  Moneda string `json:"moneda" binding:"required"`
  Monto float64 `json:"monto" binding:"required"`
  Emisor string `json:"emisor" binding:"required"`
  Receptor string `json:"receptor" binding:"required"`
  Fecha string `json:"date_transaction" binding:"required"`
}
