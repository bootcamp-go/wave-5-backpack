package tickets

import (
	"desafio_web/internal/domain"
	"desafio_web/pkg/storage"
	"fmt"
)

type Repository interface {
	GetTicketsByCountry(destination string) ([]domain.Ticket, error)
	GetAll() ([]domain.Ticket, error)
}

type repository struct {
	db storage.Store
}

func NewRepository(db storage.Store) Repository {
	return &repository{db: db}
}

func (r *repository) GetTicketsByCountry(destination string) ([]domain.Ticket, error) {
	var tickets []domain.Ticket
	var tbyc []domain.Ticket
	err := r.db.Read(&tickets)
	if err != nil {
		return nil, fmt.Errorf("Ocurrió un error al intentar leer el archivo")
	}
	for _, t := range tickets {
		if t.Country == destination {
			tbyc = append(tbyc, t)
		}
	}
	if len(tickets) == 0 {
		return nil, fmt.Errorf("No hay ningún registro")
	}
	return tbyc, nil
}
func (r *repository) GetAll() ([]domain.Ticket, error) {
	var tickets []domain.Ticket
	err := r.db.Read(&tickets)
	if err != nil {
		return nil, fmt.Errorf("Ocurrió un error al intentar leer el archivo")
	}
	if len(tickets) == 0 {
		return nil, fmt.Errorf("No hay ningún registro")
	}
	return tickets, nil
}
