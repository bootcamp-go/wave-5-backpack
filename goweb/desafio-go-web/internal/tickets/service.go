package tickets

import (
	"context"
	"desafio-go-web/internal/domain"
	"fmt"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
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

func (s service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't get tickets: %w", err)
	}
	return tickets, nil
}

func (s service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return nil, fmt.Errorf("couldn't get tickets: %w", err)
	}
	return tickets, nil
}

func (s service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	tickets, err := s.GetAll(ctx)
	if err != nil {
		return 0, fmt.Errorf("couldn't get tickets: %w", err)
	}
	count := 0.0
	for i := range tickets {
		if tickets[i].Country == destination {
			count++
		}
	}
	return count / float64(len(tickets)), nil
}
