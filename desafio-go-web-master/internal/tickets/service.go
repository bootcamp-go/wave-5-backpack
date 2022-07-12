package tickets

import (
	"context"
	"desafio-go-web/internal/domain"
	"fmt"
)

type Service interface {
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (int, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	totalTickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return []domain.Ticket{}, err
	}

	return totalTickets, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (int, error) {
	ticketDestination, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}

	totalTickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}

	avg := len(ticketDestination) / len(totalTickets)

	return avg, nil
}
