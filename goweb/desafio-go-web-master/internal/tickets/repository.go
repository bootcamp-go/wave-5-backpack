package tickets

import (
	"context"
	"desafio-go-web/internal/domain"
	"fmt"
)

type Repository interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
}

type repository struct {
	db []domain.Ticket
}

func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Ticket, error) {
	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}
	return r.db, nil
}

func (r *repository) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}
	var ticketsDest []domain.Ticket
	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}
	return ticketsDest, nil
}
