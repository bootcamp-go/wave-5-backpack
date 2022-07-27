package tickets

import (
	"context"
	"desafio-go-web/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, dest string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, dest string) (int, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	allTk, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return allTk, nil
}

func (s *service) GetTicketByDestination(ctx context.Context, dest string) ([]domain.Ticket, error) {
	tk, err := s.repository.GetTicketByDestination(ctx, dest)
	if err != nil {
		return nil, err
	}
	return tk, nil
}

func (s *service) AverageDestination(ctx context.Context, dest string) (int, error) {
	allTk, err := s.GetAll(ctx)
	if err != nil {
		return 0, err
	}
	tkByDestination, err := s.GetTicketByDestination(ctx, dest)
	if err != nil {
		return 0, err
	}
	return tkByDestination/len(allTk), nil
}