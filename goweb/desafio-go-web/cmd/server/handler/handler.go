package handler

import (
	"net/http"

	"desafio-go-web/internal/tickets"
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

func (t *Ticket) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := t.service.GetTotalTickets(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (t *Ticket) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := t.service.AverageDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}

func (t *Ticket) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		tickets, err := t.service.GetAll(c)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}
