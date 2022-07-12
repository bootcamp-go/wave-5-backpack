package tickets

import (
	"context"
	"errors"
)

type Service interface {
	AverageDestination(ctx context.Context, destination string) (float64, error)
	GetTotalTickets(ctx context.Context, destination string) (int, error)
}

type service struct {
	r Repository
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	ticketsTotal, err := s.r.GetAll(ctx)

	if err != nil {
		return 0, errors.New("Ocurrió un error al buscar todos los tickets.")
	}

	ticketsByDestino, err := s.r.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, errors.New("Ocurrió un error al buscar los tickets por destino.")
	}

	var avg float64

	if len(ticketsByDestino) == 0 || len(ticketsTotal) == 0 {
		return 0, errors.New("No hay registros.")
	}
	avg = float64(len(ticketsByDestino)) * float64(100) / float64(len(ticketsTotal))

	return float64(avg), nil
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	ticketsTotal, err := s.r.GetTicketByDestination(ctx, destination)

	if err != nil {
		return 0, errors.New("Ocurrió un error al buscar todos los tickets.")
	}

	return len(ticketsTotal), nil
}

func NewService(r Repository) Service {
	return &service{
		r: r,
	}
}
