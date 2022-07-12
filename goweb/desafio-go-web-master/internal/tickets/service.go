package tickets

import (
	"context"
)

type service struct {
	repository Repository
}

type Service interface {
	GetTotalTickets(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	tickets_country, err := s.GetTotalTickets(ctx, destination)
	if err != nil {
		return 0, err
	}
	tickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return 0, err
	}
	average := float64(tickets_country) / float64(len(tickets))
	return average, err
}
