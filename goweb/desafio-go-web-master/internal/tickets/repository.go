package tickets

import (
	"desafio-go-web/internal/domain"
	"desafio-go-web/pkg/store"
	"errors"
)

type Repository interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	GetAverageByDestination(destination string) (int, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Ticket, error) {
	lista, err := r.db.LoadTicketsFromFile()
	if err != nil {
		return nil, errors.New("could not get tickets from source")
	}
	return lista, nil
}

func (r *repository) GetTicketByDestination(destination string) ([]domain.Ticket, error) {
	lista, err := r.db.LoadTicketsFromFile()
	if err != nil {
		return nil, errors.New("could not get tickets from source")
	}
	if len(lista) == 0 {
		return nil, errors.New("empty list of tickets")
	}

	var ticketsDest []domain.Ticket
	for _, t := range lista {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}

func (r *repository) GetAverageByDestination(destination string) (int, error) {
	lista, err := r.db.LoadTicketsFromFile()
	if err != nil {
		return 0, errors.New("could not get tickets from source")
	}
	if len(lista) == 0 {
		return 0, errors.New("empty list of tickets")
	}
	var ticketsDest int
	for _, t := range lista {
		if t.Country == destination {
			ticketsDest++
		}
	}

	return ((ticketsDest * 100) / len(lista)), nil
}
