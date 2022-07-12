package tickets

import (
	"context"
)

type Service interface {
	GetTotalTickets(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
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
	ticketsTotal, err := s.repository.GetAll(ctx)
	if err != nil {
		return 0, err
	}

	ticketsDestination, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}

	return float64(len(ticketsDestination)) / float64(len(ticketsTotal)), nil
}
