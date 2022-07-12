package tickets

import (
	"desafio-go-web/internal/domain"
	"fmt"
)

type Service interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	AverageDestination(destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Ticket, error) {
	tickets, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return tickets, nil

}

func (s *service) GetTicketByDestination(destination string) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (s *service) AverageDestination(destination string) (float64, error) {

	ticket, err := s.repository.GetTicketByDestination(destination)
	fmt.Println("---------->Ticket", len(ticket))
	if err != nil {
		return 0, fmt.Errorf("error al obtener tickets con el destino: %s", destination)
	}

	if len(ticket) == 0 {
		return 0, fmt.Errorf("empty list of tickets")
	}

	tickets, err := s.repository.GetAll()
	fmt.Println("---------->Ticket", len(tickets))
	if err != nil {
		return 0, err
	}
	response := float64(len(ticket)) / float64(len(tickets)) 
	return response, nil
}
