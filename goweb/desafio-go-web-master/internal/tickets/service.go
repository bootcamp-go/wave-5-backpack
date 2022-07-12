package tickets

import (
	//"desafio-go-web/internal/domain"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetTotalTickets(c *gin.Context, destination string) (int, error)
	AverageDestination(c *gin.Context, destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(c *gin.Context, destination string) (int, error) {
	ps, err := s.repository.GetTicketByDestination(c, destination)
	if err != nil {
		return 0, err
	}
	return len(ps), nil
}

func (s *service) AverageDestination(c *gin.Context, destination string) (float64, error) {

	cantDestiny, err := s.repository.GetTicketByDestination(c, destination)
	cantTotal, err2 := s.repository.GetAll(c)

	if err != nil {
		return 0, err
	}
	if err2 != nil {
		return 0, err
	}

	ave := float64(len(cantDestiny)) / float64(len(cantTotal))
	println(len(cantDestiny))
	println(len(cantTotal))
	println(ave)
	return ave, nil

}
