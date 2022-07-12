package tickets

import (
	"context"
	"desafio-go-web/internal/domain"
)

type Service interface {
	GetAverageByDestination(ctx context.Context, destination string) (float32, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAverageByDestination(ctx context.Context, destination string) (float32, error) {
	
	totalTickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return 0, err
	}

	ticketsByDestination, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}

	average := float64(len(ticketsByDestination)) / float64(len(totalTickets)) * 100.0
	if average == 0 {
		return 0, err
	}
	return float32(average), err

}

func (s *service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	ticketsByDestination, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return nil, err
	}

	return ticketsByDestination, err
}