package handler

import (
	"net/http"

	tickets "desafio-go-web/internal/ticket"

	"github.com/gin-gonic/gin"
)

type ticket struct {
	service tickets.Service
}

func NewTicket(s tickets.Service) *ticket {
	return &ticket{
		service: s,
	}
}

func (s *ticket) GetTicketsByCountry() gin.HandlerFunc {
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

func (s *ticket) AverageDestination() gin.HandlerFunc {
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
