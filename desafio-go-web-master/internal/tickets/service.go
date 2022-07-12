package tickets

import (
	"github.com/bootcamp-go/wave-5-backpack/desafio-go-web-master/internal/domain"
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
	ticks, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ticks, nil
}

func (s *service) GetTotalTickets(destination string) ([]domain.Ticket, error) {
	ticketsAlDestino, eror := s.repository.GetTicketByDestination(destination)
	if eror != nil {
		return nil, eror
	}

	return ticketsAlDestino, nil
}
func (s *service) AverageDestination(destination string) (float64, error) {
	ticketsTotales, eror := s.repository.GetAll()
	if eror != nil {
		return 0, nil
	}
	ticketsADestino, eror := s.GetTotalTickets(destination)
	if eror != nil {
		return 0, nil
	}
	totales := len(ticketsTotales)
	totalADestino := len(ticketsADestino)
	return float64(totalADestino) / float64(totales) * 100, nil
}
