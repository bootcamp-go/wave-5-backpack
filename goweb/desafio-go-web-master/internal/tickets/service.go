package tickets

import (
	"goweb/desafio-go-web-master/internal/domain"
)

type Service interface {
	GetTicketsByCountry(dest string) ([]domain.Ticket, error)
	AverageDestination(dest string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTicketsByCountry(destination string) ([]domain.Ticket, error) {
	ts, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (s *service) AverageDestination(destination string) (float64, error) {
	ts, err := s.repository.GetAll()
	if err != nil {
		return 0, err
	}
	var contadorTickets = 0.0
	var contadorTicketsDest = 0.0
	for _, tickt := range ts {
		contadorTickets++
		if tickt.Country == destination {
			contadorTicketsDest++
		}
	}
	var promedio = contadorTicketsDest / contadorTickets
	return promedio, nil
}
