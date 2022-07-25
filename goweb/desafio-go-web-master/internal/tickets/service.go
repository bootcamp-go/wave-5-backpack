package tickets

import (
	"desafio-go-web/internal/domain"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetTotalTickets(c *gin.Context, destination string) ([]domain.Ticket, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(c *gin.Context, destination string) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetTicketByDestination(c, destination)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}
