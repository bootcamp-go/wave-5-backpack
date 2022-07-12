package tickets

import (
	"context"
	"desafio-go-web/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	GetTicketDestinationNum(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]domain.Ticket, error) {
	tickets, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (s *service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	tickets, err := s.repo.GetTicketByDestination(ctx, destination)
	if err != nil {
		return []domain.Ticket{}, err
	}
	return tickets, nil
}

func (s *service) GetTicketDestinationNum(ctx context.Context, destination string) (int, error) {
	tickets, err := s.repo.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	num := len(tickets)
	return num, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	tickets, err := s.repo.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	num := len(tickets)
	average := (float64(num) / 100.0)

	return average, nil
}
