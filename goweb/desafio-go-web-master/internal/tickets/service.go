package tickets

import (
	"context"
	"desafio-go-web/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(destination string) (int, error)
}

type service struct {
	rep Repository
}

func NewService(r Repository) Service {
	return &service{rep: r}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	t, err := s.rep.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	t, err := s.rep.GetTicketByDestination(ctx, destination)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (s *service) AverageDestination(destination string) (int, error) {
	a, err := s.rep.AverageDestination(destination)
	if err != nil {
		return 0, err
	}
	return a, nil
}
