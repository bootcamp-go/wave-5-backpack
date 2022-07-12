package tickets

import (
	"context"
	"desafio-go-web/internal/domain"
	"fmt"
)

//

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
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

func (r *repository) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error) {

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

func (r *repository) AverageDestination(ctx context.Context, destination string) (float64, error) {
	var sumaTotal float64
	var sumCountry float64

	for _, t := range r.db {
		sumaTotal = t.Price + sumaTotal
		if t.Country == destination {
			sumCountry = t.Price + sumCountry
		}
	}
	fmt.Println(sumaTotal)

	average := (sumCountry / sumaTotal)
	return average, nil
}
