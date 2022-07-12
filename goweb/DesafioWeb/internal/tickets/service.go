package tickets

import (
	"desafio_web/internal/domain"
	"fmt"
)

type Service interface {
	GetTicketsByCountry(destination string) ([]domain.Ticket, error)
	GetAverageCountry(destination string) (float64, error)
}

type service struct {
	rep Repository
}

func NewService(r Repository) Service {
	return &service{rep: r}
}

func (s *service) GetTicketsByCountry(destination string) ([]domain.Ticket, error) {
	tickers, err := s.rep.GetTicketsByCountry(destination)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return tickers, nil
}
func (s *service) GetAverageCountry(destination string) (float64, error) {
	allTickers, err := s.rep.GetAll()
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	conTickers, err := s.rep.GetTicketsByCountry(destination)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	average := float64(len(conTickers)) * 100 / float64(len(allTickers))
	return average, nil
}
