package tickets

import (
	"desafio-go-web/internal/domain"
	"context"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) (int, error)  {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error)  {
	ticketsEnTotal, err := s.repository.GetAll(ctx)
	if err != nil {
		return 0, err
	}

	ticketsPorDestino, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}

	resultado := float64(len(ticketsPorDestino)) / float64(len(ticketsEnTotal)) * 100
	return resultado, nil
	
}


