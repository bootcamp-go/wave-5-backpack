package tickets

import (
	"context"
	"fmt"
)

type Service interface {
	GetTotalofTickets(ctx context.Context, destination string) (int, error)
	AVGDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalofTickets(ctx context.Context, destination string) (int, error) {
	tickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return 0, fmt.Errorf("An error ocurred getting all the tickets")
	}
	var result = 0
	for _, ticket := range tickets {
		if ticket.Country == destination {
			result++
		}
	}
	if result == 0 {
		return 0, fmt.Errorf("Country not found")
	}
	return result, nil
}

func (s *service) AVGDestination(ctx context.Context, destination string) (float64, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, fmt.Errorf("An error has ocurred gettin the tickets by destination")
	}

	allTickets, err := s.repository.GetAll(ctx)

	if err != nil {
		return 0, fmt.Errorf("An error has ocurred getting all tickets")
	}

	var result float64
	if len(allTickets) != 0 {
		result = (float64(len(tickets)) * 100) / float64(len(allTickets))
	}
	return result, nil
}
