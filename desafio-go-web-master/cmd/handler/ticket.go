package handler

import (
	"fmt"
	"net/http"

	"desafio-go-web-master/internal/tickets"

	"github.com/gin-gonic/gin"
)

type Service struct {
	service tickets.Service
}

func NewTicket(s tickets.Service) *Service {
	return &Service{
		service: s,
	}
}

func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTicketByDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *Service) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {
		var avg float64
		destination := c.Param("dest")

		totalDestination, err := s.service.GetTicketByDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		totalTickets, err := s.service.GetAll(c)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		avg = (float64(len(totalDestination)) / float64(len(totalTickets))) * 100

		fmt.Println(avg)
		c.JSON(200, avg)
	}
}
