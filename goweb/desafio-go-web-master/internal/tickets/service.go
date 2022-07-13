package tickets

import "desafio-go-web/internal/domain"

type Service interface {
	GetTotalTickets() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
	AverageDestination(destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetTotalTickets() ([]domain.Ticket, error) {
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
	avg, err := s.repository.AverageDestination(destination)
	if err != nil {
		return 0.0, err
	}

	return avg, nil
}
