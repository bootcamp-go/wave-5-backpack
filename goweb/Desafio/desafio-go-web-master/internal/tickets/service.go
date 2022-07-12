package tickets

import (
	"desafio-go-web/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	AverageDestination(destination string) ([]domain.Ticket, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Ticket, error) {
	tickets, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return tickets, nil

}

func (s *service) GetTicketByDestination(destination string) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (s *service) AverageDestination(destination string) ([]domain.Ticket, error) {

	return []domain.Ticket{}, nil
}
