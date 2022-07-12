package tickets

import (
	"desafio-go-web/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]domain.Ticket, error) {
	tickets, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (s *service) GetTicketByDestination(destination string) ([]domain.Ticket, error) {
	tickets, err := s.repo.GetTicketByDestination(destination)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}
