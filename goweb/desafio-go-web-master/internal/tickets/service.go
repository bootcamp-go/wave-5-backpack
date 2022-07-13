package tickets

import (
	"context"
)

type Service interface {
	//GetAll(ctx context.Context) ([]domain.Ticket, error)
	//GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return service{
		repository: r,
	}
}

func (s service) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)

	if err != nil {
		return 0, err
	}

	totalTicketsDestination := len(tickets) - 1

	return totalTicketsDestination, nil
}

func (s service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)

	if err != nil {
		return 0, err
	}

	totalTicketsDestination := len(tickets) - 1

	tickets, err = s.repository.GetAll(ctx)

	if err != nil {
		return 0, err
	}

	totalTickets := len(tickets) - 1

	return (float64(totalTicketsDestination) / float64(totalTickets)), nil
}
