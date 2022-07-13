package tickets

import (
	"context"
	"desafio-go-web-master/internal/domain"
	"log"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
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

func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	ts, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return ts, nil
}

func (s *service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	var ticketsDest []domain.Ticket

	ticketsDest, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		log.Fatal(err)
	}

	return ticketsDest, nil
}
