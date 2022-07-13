package tickets

import (
	"github.com/gin-gonic/gin"
)

type service struct {
	r Repository
}
type Service interface {
	GetTotalTickets(*gin.Context, string) (int, error)
	AverageDestination(*gin.Context, string) (float64, error)
}

func (s *service) GetTotalTickets(ctx *gin.Context, destination string) (int, error) {
	ticketsWithDest, err := s.r.GetTicketByDestination(ctx, destination)
	return len(ticketsWithDest), err
}
func (s *service) AverageDestination(ctx *gin.Context, dest string) (float64, error) {
	listTicket, err := s.r.GetAll(ctx)
	if err != nil {
		return -1, err
	}
	totalListTicket := len(listTicket)
	totalTicketsWithDest, err := s.GetTotalTickets(ctx, dest)
	if err != nil {
		return -1, err
	}
	return float64(totalTicketsWithDest) * 100 / float64(totalListTicket), nil
}
func NewService(r Repository) Service {
	return &service{r}
}
