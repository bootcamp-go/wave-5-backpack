package tickets

import (
	"context"
	"desafio-go-web/internal/domain"
)

type Service interface {
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error)
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

func (s *service) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {

	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {

	tickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return 0, err
	}

	ticketsDestination, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}

	average := (float64(len(ticketsDestination)) / float64(len(tickets))) * 100

	return float64(average), nil
}
