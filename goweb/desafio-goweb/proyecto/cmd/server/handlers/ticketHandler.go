package handlers

import (
	"net/http"
	"proyecto/internal/ticket"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	service ticket.Service
}

func NewTicketHandler(s ticket.Service) *TicketHandler {
	return &TicketHandler{service: s}
}

func (t *TicketHandler) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		destination := c.Param("dest")

		tickets, err := t.service.GetTicketByDestination(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, len(tickets))
	}
}

func (s *TicketHandler) AverageDestination() gin.HandlerFunc {
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
