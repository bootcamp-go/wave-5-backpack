package tickets

import (
	"context"
	"errors"
)

type Service interface {
	GetTotalTickets(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	tickets, err := s.repo.GetAll(ctx)
	if err != nil {
		return 0, errors.New("error: failed loading tickets")
	}

	var result = 0
	for _, ticket := range tickets {
		if ticket.Country == destination {
			result++
		}
	}

	return result, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	tickets, err := s.repo.GetTicketByDestination(ctx, destination)

	if err != nil {
		return 0, errors.New("error: failed loading tickets by destination")
	}

	allTickets, err := s.repo.GetAll(ctx)

	if err != nil {
		return 0, errors.New("error: failed getting all tickets")
	}

	var average float64 = 0
	if len(allTickets) != 0 {
		average = (float64(len(tickets)) * 100) / float64(len(allTickets))
	}

	return average, nil
}
