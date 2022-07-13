package tickets

import (
	"desafio-go-web/internal/domain"
	"fmt"
)

type Repository interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	AverageDestination(destination string) (float64, error)
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

func (r *repository) GetTicketByDestination(destination string) ([]domain.Ticket, error) {
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

func (r *repository) AverageDestination(destination string) (float64, error) {
	var tDestination []domain.Ticket

	if len(r.db) == 0 {
		return 0.0, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			tDestination = append(tDestination, t)
		}
	}

	// return float64(len(tDestination)), nil
	// return float64(len(r.db)), nil
	return (float64(len(tDestination)) / float64(len(r.db))) * 100, nil
}
