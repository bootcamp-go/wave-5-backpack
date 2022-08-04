package handler

import (
	"fmt"
	"net/http"

	"github.com/bootcamp-go/wave-5-backpack/desafio-go-web-master/internal/tickets"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name    string  `json:"name" binding:"required"`
	Email   string  `json:"email" binding:"required"`
	Country string  `json:"country" binding:"required"`
	Time    string  `json:"time" binding:"required"`
	Price   float64 `json:"price" binding:"required"`
}

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

		message := fmt.Sprintf("Cantidad de tickets a %s: %d", destination, tickets)

		c.JSON(200, message)
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

		message := fmt.Sprintf("Promedio de personas que viajan a %s: %.2f", destination, avg)

		c.JSON(200, message)
	}
}
