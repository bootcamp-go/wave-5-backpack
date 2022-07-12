package handler

import (
	"net/http"

	"github.com/bootcamp-go/wave-5-backpack/desafio-go-web-master/internal/tickets"
	"github.com/gin-gonic/gin"
)

// type request struct {
// 	Name    string  `json:"nombre" binding:"required"`
// 	Email   string  `json:"email" binding:"required"`
// 	Country string  `json:"destino" binding:"required"`
// 	Time    string  `json:"tiempo" `
// 	Price   float64 `json:"precio" `
// }

type Ticket struct {
	service tickets.Service
}

func NewTicket(s tickets.Service) *Ticket {
	return &Ticket{
		service: s,
	}
}

func (s *Ticket) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTotalTickets(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *Ticket) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ticketsTotales, err := s.service.GetAll()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}
		c.JSON(200, ticketsTotales)
	}
}

func (s *Ticket) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.AverageDestination(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}
