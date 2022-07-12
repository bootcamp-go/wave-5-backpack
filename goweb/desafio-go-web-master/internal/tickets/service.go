package tickets

import (
	"context"
)

type Service interface {
	GetTotalTickets(c context.Context, destination string) (int, error)
	AverageDestination(c context.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetTotalTickets(c context.Context, destination string) (int, error) {
	tickets, err := s.repository.GetTicketByDestination(c, destination)
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}

func (s *service) AverageDestination(c context.Context, destination string) (float64, error) {
	tickets, err := s.repository.GetAll(c)
	if err != nil {
		return 0, err
	}
	cantidad, err := s.GetTotalTickets(c, destination)
	if err != nil {
		return 0, err
	}
	avg := float64(cantidad) / float64(len(tickets))
	return avg, nil
}
