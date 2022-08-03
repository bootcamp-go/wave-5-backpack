package transactions

import "log"

type Service interface {
	GetAll() ([]*Transaction, error)
	Store(codigo, moneda, emisor, receptor string, monto float64) (*Transaction, error)
	Update(id int64, codigo, moneda, emisor, receptor string, monto float64) (*Transaction, error)
	UpdateReceptorYMonto(id int64, receptor string, monto float64) (*Transaction, error)
	Delete(id int64) error
}

type service struct {
	rep Repository
}

func NewService(r Repository) Service {
	return &service{rep: r}
}

func (s *service) GetAll() ([]*Transaction, error) {
	return s.rep.GetAll()
}

func (s *service) Store(codigo, moneda, emisor, receptor string, monto float64) (*Transaction, error) {
	lastId, err := s.rep.LastId()
	if err != nil {
		return nil, err
	}

	lastId++
	log.Println("Se creara un id: ", lastId)
	return s.rep.Store(lastId, codigo, moneda, emisor, receptor, monto)
}

func (s *service) Update(id int64, codigo, moneda, emisor, receptor string, monto float64) (*Transaction, error) {
	return s.rep.Update(id, codigo, moneda, emisor, receptor, monto)
}

func (s *service) UpdateReceptorYMonto(id int64, receptor string, monto float64) (*Transaction, error) {
	return s.rep.UpdateReceptorYMonto(id, receptor, monto)
}

func (s *service) Delete(id int64) error {
	return s.rep.Delete(id)
}
