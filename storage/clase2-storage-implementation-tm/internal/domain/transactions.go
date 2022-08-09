package domain

type Transaction struct {
	ID                int     `json:"id" binding:"-"`
	CodigoTransaccion string  `json:"codigo de transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	Fecha             string  `json:"fecha de transaccion"`
}
