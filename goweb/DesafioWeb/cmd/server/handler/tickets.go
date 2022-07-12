package handler

import (
	"desafio_web/internal/tickets"
	"net/http"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Country string  `json:"country"`
	Time    string  `json:"time"`
	Price   float64 `json:"price"`
}
type Service struct {
	service tickets.Service
}

func NewUser(s tickets.Service) *Service {
	return &Service{service: s}
}
func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTicketsByCountry(destination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, tickets)
	}
}

func (s *Service) GetAverageCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.GetAverageCountry(destination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, avg)
	}
}
