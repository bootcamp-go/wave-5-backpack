package tickets

import (
	"github.com/bootcamp-go/wave-5-backpack/goweb_desafio/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Ticket, error)
	GetTotalTickets(destination string) ([]domain.Ticket, error)
	AverageDestination(destination string) (float64, error)
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Ticket, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) GetTotalTickets(destination string) ([]domain.Ticket, error) {
	ps, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) AverageDestination(destination string) (float64, error) {
	totalTickets, err := s.repository.GetAll()
	if err != nil {
		return 0, err
	}

	destinationTickets, err2 := s.repository.GetTicketByDestination(destination)
	if err2 != nil {
		return 0, err
	}
	total := float64(len(totalTickets))
	dest := float64(len(destinationTickets))
	result := (dest / total) * 100

	return result, nil
}
