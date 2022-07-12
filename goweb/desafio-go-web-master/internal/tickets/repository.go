package tickets

import (
	"desafio-go-web-master/internal/domain"
	"fmt"
)

type Repository interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	//GetTotalTickets(destination string) ([]domain.Ticket, error)
	//AverageDestination(destination string) ([]domain.Ticket, error)
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

	//fmt.Println(r.db)
	if len(r.db) == 0 {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range r.db {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	//fmt.Println(ticketsDest)
	return ticketsDest, nil
}

// func (r *repository) GetTotalTickets(destination string) ([]domain.Ticket, error) {
// 	var ticketsDest []domain.Ticket

// 	if len(r.db) == 0 {
// 		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
// 	}

// 	for _, t := range r.db {
// 		if t.Country == destination {
// 			ticketsDest = append(ticketsDest, t)
// 		}
// 	}

// 	return ticketsDest, nil
// }

// func (r *repository) AverageDestination(destination string) ([]domain.Ticket, error) {
// 	var ticketsDest []domain.Ticket

// 	if len(r.db) == 0 {
// 		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
// 	}

// 	for _, t := range r.db {
// 		if t.Country == destination {
// 			ticketsDest = append(ticketsDest, t)
// 		}
// 	}

// 	return ticketsDest, nil
// }
