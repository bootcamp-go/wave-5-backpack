package transactions

type Service interface {
	GetAll() ([]Transaction, error)
	Store(codigo, moneda, emisor, receptor string, monto float64) (Transaction, error)
	Update(id int64, monto float64, codigo, emisor, receptor, moneda string) (Transaction, error)
	UpdateReceptorYMonto(id int64, receptor string, monto float64) (Transaction, error)
	Delete(id int64) error
}

type service struct {
	rep Repository
}

func NewService(r Repository) Service {
	return &service{rep: r}
}

func (s *service) GetAll() ([]Transaction, error) {
	transactions, err := s.rep.GetAll()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *service) Store(codigo, moneda, emisor, receptor string, monto float64) (Transaction, error) {
	lastId, err := s.rep.LastId()
	if err != nil {
		return Transaction{}, err
	}

	lastId++
	transaction, err := s.rep.Store(lastId, codigo, moneda, emisor, receptor, monto)
	if err != nil {
		return Transaction{}, err
	}
	return transaction, nil
}

func (s *service) Update(id int64, monto float64, codigo, emisor, receptor, moneda string) (Transaction, error) {
	return s.rep.Update(id, monto, codigo, emisor, receptor, moneda)
}

func (s *service) Delete(id int64) error {
	return s.rep.Delete(id)
}
func (s *service) UpdateReceptorYMonto(id int64, receptor string, monto float64) (Transaction, error) {
	return s.rep.UpdateReceptorYMonto(id, receptor, monto)
}
