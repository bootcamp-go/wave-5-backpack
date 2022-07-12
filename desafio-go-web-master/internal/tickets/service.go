package tickets

import (
	"desafio-go-web/internal/domain"
)

type Service interface {
	GetTotalTickets(destination string) ([]domain.Ticket, error)
	AverageDestination(destination string) (float64, error)
}

type service struct {
	rep Repository
}

func NewService(r Repository) Service {
	return &service{rep: r}
}

func (s *service) GetTotalTickets(destination string) ([]domain.Ticket, error) {
	return s.rep.GetTicketByDestination(destination)
}

func (s *service) AverageDestination(destination string) (float64, error) {
	var average float64

	ticketsByCountry, err := s.rep.GetTicketByDestination(destination)

	if err != nil {
		return 0, err
	}

	totalTickets, err := s.rep.GetAll()

	if err != nil {
		return 0, err
	}

	average = float64(len(totalTickets)) / float64(len(ticketsByCountry))

	return average, nil
}
