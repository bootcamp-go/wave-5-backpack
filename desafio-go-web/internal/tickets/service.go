package tickets

import (
	"desafio-go-web/internal/domain"
	"fmt"
)

type Service interface {
	AverageDestination(destination string) (float64, error)
	GetDestination(destination string) ([]domain.Ticket, error)
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

type service struct {
	repository Repository
}

func (s *service) GetDestination(destination string) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return []domain.Ticket{}, fmt.Errorf("error al obtener Tickets de %s. error: %s", destination, err)
	}

	return tickets, nil
}

func (s *service) AverageDestination(destination string) (float64, error) {
	ticketDestination, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return 0, fmt.Errorf("error al obtener Tickets de %s. Error: %s", destination, err)
	}

	if len(ticketDestination) == 0 {
		return 0, fmt.Errorf("no existen tickets para %s. Imposible sacar promedio", destination)
	}

	tickets, err := s.repository.GetAll()
	if err != nil {
		return 0, fmt.Errorf("error al obtener tickets: %s", err)
	}

	avg := float64(len(ticketDestination)) / float64(len(tickets))

	return avg, nil
}
