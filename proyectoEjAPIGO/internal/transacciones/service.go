package transacciones

import (
	"fmt"
	"goweb/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Transaccion, error)
	Store(codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (domain.Transaccion, error)
	Update(id int, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (domain.Transaccion, error)
	UpdateCTandMonto(id int, codigo_transaccion string, monto float64) (domain.Transaccion, error)
	Delete(id int) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]domain.Transaccion, error) {
	transactions, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *service) Store(codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (domain.Transaccion, error) {
	lastID, err := s.repo.LastID()
	if err != nil {
		return domain.Transaccion{}, fmt.Errorf("error al traer el ultimno id %w", err)
	}
	lastID++

	transaction, err := s.repo.Store(lastID, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion, monto)
	if err != nil {
		return domain.Transaccion{}, fmt.Errorf("error al crear la transaccion: %w", err)
	}
	return transaction, nil
}

func (s *service) Update(id int, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (domain.Transaccion, error) {
	t, err := s.repo.Update(id, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion, monto)
	if err != nil {
		return domain.Transaccion{}, fmt.Errorf("errror al actualizar la transaccion %w", err)
	}
	return t, nil
}

func (s *service) UpdateCTandMonto(id int, codigo_transaccion string, monto float64) (domain.Transaccion, error) {
	t, err := s.repo.UpdateCTandMonto(id, codigo_transaccion, monto)
	if err != nil {
		return domain.Transaccion{}, fmt.Errorf("error al actualizar la transaccion %w", err)
	}
	return t, nil
}

func (s *service) Delete(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("error al eleminar la transaccion %w", err)
	}
	return nil
}
