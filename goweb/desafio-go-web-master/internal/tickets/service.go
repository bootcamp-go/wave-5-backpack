package tickets

import (
	"context"
)

type Service interface {
	GetTotalTickets(context.Context, string) (float64, error)
	AverageDestination(context.Context, string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) (float64, error) {
	var totalPrice float64

	ticketsDest, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}

	for _, i := range ticketsDest {
		totalPrice += i.Price
	}
	return totalPrice, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {

	ticketsDest, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}

	totalTickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return 0, err
	}

	averageTickets := float64(len(ticketsDest)) / float64(len(totalTickets)) * 100

	return averageTickets, nil
}
