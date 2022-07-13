package tickets

import (
	"context"
	"desafio-go-web/internal/domain"
	"fmt"
)

type Service interface {
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (string, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (string, error) {
	tickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return "", err
	}

	totalDestination := 0
	for i := range tickets {
		if tickets[i].Country == destination {
			totalDestination++
		}
	}

	average := float64(totalDestination) * (100.0 / float64(len(tickets)))

	return fmt.Sprintf("El promedio de personas que viajan a %s por día son: %.2f%%", destination, average), nil
}
