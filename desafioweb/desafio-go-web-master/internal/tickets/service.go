package tickets

import (
	"desafio-go-web-master/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Ticket, error)
	GetTotalTickets(country string) ([]domain.Ticket, error)
	AverageDestination(country string) (float64, error)
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
	tickets, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (s *service) GetTotalTickets(country string) ([]domain.Ticket, error) {
	ticketsDest, err := s.repository.GetTicketByDestination(country)
	if err != nil {
		return nil, err
	}

	return ticketsDest, nil
}

func (s *service) AverageDestination(country string) (float64, error) {
	allTickets, err := s.repository.GetAll()
	if err != nil {
		return 0, err
	}

	ticketsDest, err := s.repository.GetTicketByDestination(country)
	if err != nil {
		return 0, err
	}

	totalTickets := len(allTickets)
	totalTicketsDest := len(ticketsDest)

	return float64(totalTicketsDest / totalTickets), nil
}
