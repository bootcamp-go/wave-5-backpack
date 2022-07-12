package handler

import (
	"desafio-go-web/interal/tickets"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type request struct {
// 	Name    string  `json:"nombre" binding:"required"`
// 	Email   string  `json:"email" binding:"required"`
// 	Country string  `json:"destino" binding:"required"`
// 	Time    string  `json:"tiempo" `
// 	Price   float64 `json:"precio" `
// }

type Service struct {
	service tickets.Service
}

func NewService(s tickets.Service) *Service {
	return &Service{
		service: s,
	}
}

func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTotalTickets(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *Service) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ticketsTotales, err := s.service.GetAll(c)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}
		c.JSON(200, ticketsTotales)
	}
}

func (s *Service) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.AverageDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}
