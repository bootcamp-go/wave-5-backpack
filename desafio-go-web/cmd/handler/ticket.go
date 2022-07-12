package handler

import (
	"net/http"

	"desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Tickets struct {
	service tickets.Service
}

func NewTickets(s tickets.Service) *Tickets {
	return &Tickets{
		service: s,
	}
}

func (s *Tickets) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetDestination(destination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": tickets})
	}
}

func (s *Tickets) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.AverageDestination(destination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": avg})
	}
}
