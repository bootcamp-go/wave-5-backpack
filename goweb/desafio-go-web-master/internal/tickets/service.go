package tickets

import (
	"desafio-go-web/internal/domain"
	"errors"
)

type Service interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketsByCountry(destination string) ([]domain.Ticket, error)
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
	tickets, err := s.repository.GetAll()
	if err != nil {
		return []domain.Ticket{}, errors.New("no existen datos")
	}
	return tickets, nil
}
func (s *service) GetTicketsByCountry(destination string) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return []domain.Ticket{}, errors.New("no existen datos")
	}
	return tickets, nil
}
func (s *service) AverageDestination(destination string) (float64, error) {
	tickets, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return 0, err
	}
	count := 0
	var suma float64 = 0
	for _, t := range tickets {
		suma += t.Price
		count++
	}
	return suma / float64(count), nil
}
