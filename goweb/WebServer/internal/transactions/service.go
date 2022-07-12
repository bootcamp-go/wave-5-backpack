package transactions

type Service interface {
	GetAll() ([]Transaction, error)
	Create(codigoTransaccion int, moneda string, monto float64, emisor, receptor, fechaTransaccion string) (Transaction, error)
	Update(id, codigoTransaccion int, moneda string, monto float64, emisor, receptor, fechaTransaccion string) (Transaction, error)
	UpdatePartial(id, codigoTransaccion int, monto float64) (Transaction, error)
	Delete(id int) (Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Transaction, error) {
	tr, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return tr, nil
}

func (s *service) Create(codigoTransaccion int, moneda string, monto float64, emisor, receptor, fechaTransaccion string) (Transaction, error) {
	lastID, err := s.repository.LastId()
	if err != nil {
		return Transaction{}, err
	}

	lastID++

	producto, err := s.repository.Create(lastID, codigoTransaccion, moneda, monto, emisor, receptor, fechaTransaccion)
	if err != nil {
		return Transaction{}, err
	}

	return producto, nil
}

func (s *service) Update(id, codigoTransaccion int, moneda string, monto float64, emisor, receptor, fechaTransaccion string) (Transaction, error) {
	return s.repository.Update(id, codigoTransaccion, moneda, monto, emisor, receptor, fechaTransaccion)
}

func (s *service) UpdatePartial(id, codigoTransaccion int, monto float64) (Transaction, error) {
	return s.repository.UpdatePartial(id, codigoTransaccion, monto)
}
func (s *service) Delete(id int) (Transaction, error) {
	return s.repository.Delete(id)
}
