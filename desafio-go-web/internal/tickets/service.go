package tickets

import (
	"desafio-go-web/internal/domain"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Service interface {
	AverageDestination(c *gin.Context, destination string) (float64, error)
	GetTotalTickets(c *gin.Context, destination string) ([]domain.Ticket, error)
}

type service struct {
	repository Repository
}

func (s *service) GetTotalTickets(c *gin.Context, destination string) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetTicketByDestination(c, destination)
	if err != nil {
		return []domain.Ticket{}, fmt.Errorf("error al obtener Tickets de %s. error: %s", destination, err)
	}

	return tickets, nil
}

func (s *service) AverageDestination(c *gin.Context, destination string) (float64, error) {
	ticketDestination, err := s.repository.GetTicketByDestination(c, destination)
	if err != nil {
		return 0, fmt.Errorf("error al obtener Tickets de %s. Error: %s", destination, err)
	}

	if len(ticketDestination) == 0 {
		return 0, fmt.Errorf("no existen tickets para %s. Imposible sacar promedio", destination)
	}

	tickets, err := s.repository.GetAll(c)
	if err != nil {
		return 0, fmt.Errorf("error al obtener tickets: %s", err)
	}

	avg := float64(len(ticketDestination)) / float64(len(tickets))

	return avg, nil
}
