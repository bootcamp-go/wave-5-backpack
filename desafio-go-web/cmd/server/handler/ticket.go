package handler

import (
	"desafio-go-web/internal/tickets"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ticket struct {
	service tickets.Service
}

func NewTicket(s tickets.Service) *Ticket {
	return &Ticket{
		service: s,
	}
}

func (s *Ticket) GetTicketsByCountry(c *gin.Context) {
	destination := c.Param("dest")

	tickets, err := s.service.GetTicketsByCountry(destination)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error(), nil)
		return
	}

	c.JSON(200, tickets)
}

func (s *Ticket) AverageDestination(c *gin.Context) {
	destination := c.Param("dest")

	avg, err := s.service.AverageDestination(destination)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error(), nil)
		return
	}

	c.JSON(200, avg)
}
