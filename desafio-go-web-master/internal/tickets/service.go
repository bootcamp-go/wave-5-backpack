package tickets

import (
	"context"
	"desafio-go-web/internal/domain"
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

func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	ts, err := s.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (s *service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	ts, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	ts, err := s.repository.AverageDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	return ts, nil
}
