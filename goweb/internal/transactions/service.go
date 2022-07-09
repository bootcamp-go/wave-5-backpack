package transactions

type Service interface {
	GetAll() ([]*Transactions, error)
	Store(code, currency, transmitter, receiver, date string, amount float64, completed bool) (*Transactions, error)
	Update(id int64, code string, currency, transmitter, receiver, date string, amount float64, completed bool) (*Transactions, error)
	UpdateTransmitter(id int64, transmitter string) (*Transactions, error)
	Delete(id int64) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]*Transactions, error) {
	transactions, error := s.repo.GetAll()
	if error != nil {
		return nil, error
	}
	return transactions, nil
}

func (s *service) Store(code, currency, transmitter, receiver, date string, amount float64, completed bool) (*Transactions, error) {
	transactions, error := s.repo.Store(code, currency, transmitter, receiver, date, amount, completed)
	if error != nil {
		return nil, error
	}
	return transactions, nil
}

func (s service) Update(id int64, code, currency, transmitter, receiver, date string, amount float64, completed bool) (*Transactions, error) {
	return s.repo.Update(id, code, currency, transmitter, receiver, date, amount, completed)
}

func (s service) UpdateTransmitter(id int64, transmitter string) (*Transactions, error) {
	return s.repo.UpdateTransmitter(id, transmitter)
}

func (s service) Delete(id int64) error {
	return s.repo.Delete(id)
}
