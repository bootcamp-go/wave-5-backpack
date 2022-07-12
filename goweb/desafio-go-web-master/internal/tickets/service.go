package tickets

import (
	"desafio-go-web/internal/domain"
	"errors"
)

type Service interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	GetAverageByDestination(destination string) (int, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Ticket, error) {
	lista, err := s.repository.GetAll()
	if err != nil {
		return nil, errors.New("could not get info in method GetAll")
	}
	return lista, nil
}

func (s *service) GetTicketByDestination(destination string) ([]domain.Ticket, error) {
	lista, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return nil, errors.New("could not get info in method GetTicketByDestination")
	}
	return lista, nil
}

func (s *service) GetAverageByDestination(destination string) (int, error) {
	average, err := s.repository.GetAverageByDestination(destination)
	if err != nil {
		return 0, errors.New("could not get info in method GetAverageByDestination")
	}
	return average, nil
}
