package domain

type Transaction struct {
	Id                int     `json:"Id"`
	CodigoTransaccion string  `json:"codigo_transaccion" binding:"required"`
	Moneda            string  `json:"moneda" binding:"required"`
	Monto             float64 `json:"monto" binding:"required"`
	Emisor            string  `json:"emisor" binding:"required"`
	Receptor          string  `json:"receptor" binding:"required"`
	FechaTransaccion  string  `json:"fecha_transaccion" binding:"required"`
}
