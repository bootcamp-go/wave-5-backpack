package tickets

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetTotalTickets(c *gin.Context, destination string) (int, error)
	AverageDestination(c *gin.Context, destination string) (float64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetTotalTickets(c *gin.Context, destination string) (int, error) {
	tickets, err := s.repo.GetTicketByDestination(c, destination)
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}

func (s *service) AverageDestination(c *gin.Context, destination string) (float64, error) {

	totalTickets, errTotal := s.repo.GetAll(c)
	if errTotal != nil {
		return 0, errTotal
	}

	tickets, errTicket := s.repo.GetTicketByDestination(c, destination)
	if errTicket != nil {
		return 0, errTicket
	}

	return 100 * float64(len(tickets)) / float64(len(totalTickets)), nil
}
