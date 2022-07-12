package ticket

import (
	"errors"
	"proyecto/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	AverageDestination(destination string) (float64, error)
}

type service struct {
	r Repository
}

func NewService(repo Repository) Service {
	return &service{r: repo}
}

func (s *service) GetAll() ([]domain.Ticket, error) {
	return s.r.GetAll()
}

func (s *service) GetTicketByDestination(destination string) ([]domain.Ticket, error) {
	tickets, err := s.r.GetTicketByDestination(destination)
	if err != nil {
		return nil, errors.New("No hay tickets cargados en la base de datos")
	}

	if len(tickets) == 0 {
		return tickets, errors.New("No se encontraron tickets con ese destino")
	}

	return tickets, nil
}

func (s *service) AverageDestination(destination string) (float64, error) {
	ticketTotales, _ := s.r.GetAll()
	ticketDestino, _ := s.r.GetTicketByDestination(destination)
	return float64(len(ticketDestino)) / float64(len(ticketTotales)), nil
}
