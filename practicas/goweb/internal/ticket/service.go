package tickets

import (
	"desafio-go-web/internal/domain"
)

type Service interface {
	GetTotalTickets(destination string) ([]domain.Ticket, error)
	AverageDestination(destination string) ([]domain.Ticket, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) AverageDestination(destination string) ([]domain.Ticket, error) {
	filteredTickets, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return []domain.Ticket{}, err
	}
	return filteredTickets, nil

}

func (s *service) GetTotalTickets(destination string) ([]domain.Ticket, error) {
	allTickets, err := s.repository.GetAll()
	if err != nil {
		return []domain.Ticket{}, err
	}
	return allTickets, nil
}
