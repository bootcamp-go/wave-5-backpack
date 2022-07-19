package tickets

import (
	"desafio-go-web-master/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	AverageDestination(destination string) (float64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]domain.Ticket, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) GetTicketByDestination(destination string) ([]domain.Ticket, error) {
	products, err := s.repo.GetTicketByDestination(destination)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) AverageDestination(destination string) (float64, error) {
	total, err := s.repo.GetAll()
	if err != nil {
		return 0, err
	}

	pais, err := s.repo.GetTicketByDestination(destination)
	if err != nil {
		return 0, err
	}
	return float64(len(total) / len(pais)), nil
}
