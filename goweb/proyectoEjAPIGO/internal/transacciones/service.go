package transacciones

type Service interface {
	getAll() ([]Transaccion, error)
}

type service struct {
	repo Repository
}
