package tickets

import (
	"context"
	"desafio-go-web/internal/domain"
	"fmt"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(destination string) (int, error)
}

type repository struct {
	db []domain.Ticket
}

func NewRepository(db []domain.Ticket) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Ticket, error) {

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	return r.db, nil
}

func (r *repository) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {

	var ticketsDest []domain.Ticket

	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}

func (r *repository) AverageDestination(destination string) (int, error) {
	var tickets []domain.Ticket

	if len(r.db) == 0 {
		return 0, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			tickets = append(tickets, t)
		}
	}
	totalTickets := len(tickets)
	dbTickets := len(r.db)
	res := totalTickets * dbTickets / 100
	return res, nil
}
