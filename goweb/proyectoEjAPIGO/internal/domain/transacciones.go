package domain

type Transaccion struct {
	ID                 int     `json:"id"`
	Codigo_transaccion string  `json:"codigo_transaccion" binding:"required"`
	Moneda             string  `json:"moneda"`
	Monto              float64 `json:"monto"`
	Emisor             string  `json:"emisor"`
	Receptor           string  `json:"receptor"`
	Fecha_transaccion  string  `json:"fecha_transaccion"`
}
