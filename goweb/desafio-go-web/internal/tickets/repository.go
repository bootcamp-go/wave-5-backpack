package tickets

import (
	"context"
	"desafio-go-web/internal/domain"
	"desafio-go-web/pkg/store"
	"fmt"
)

const (
	errorLectura   = "no se puede leer la db, error: %s"
	errorEscritura = "no se puede escribir en la db, error: %s"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	// Leemos el CSV
	var tickets []domain.Ticket
	t, err := r.db.Read(&tickets)
	if err != nil {
		return nil, fmt.Errorf(errorLectura, err)
	}

	if len(t) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return t, nil
}

func (r *repository) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	// Leemos el CSV
	var ticketsDest []domain.Ticket
	t, err := r.db.Read(&ticketsDest)
	if err != nil {
		return nil, fmt.Errorf(errorLectura, err)
	}

	if len(t) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range t {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}
