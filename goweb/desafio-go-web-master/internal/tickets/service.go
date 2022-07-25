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
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {
	destinity, err := s.repo.GetTicketByDestination(ctx, destination)

	if err != nil {
		return nil, err
	}
	return destinity, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	totalTickets, err := s.repo.GetAll(ctx)

	if err != nil {
		return 0, err
	}

	ticketsByDestination, err := s.repo.GetTicketByDestination(ctx, destination)

	if err != nil {
		return 0, err
	}

	average := (float64(len(ticketsByDestination)) / float64(len(totalTickets))) * 100
	return float64(average), nil
}
