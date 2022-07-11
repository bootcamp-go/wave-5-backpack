package transacciones

import 
type Service interface {
	getAll() ([]transacciones.Transaccion, error)
	Store(id int, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (Transaccion, error)
	Update(id int, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (Transaccion, error)
	UpdateCTandMonto(id int, codigo_transaccion string, monto float64) (Transaccion, error)
	Delete(id int) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}
