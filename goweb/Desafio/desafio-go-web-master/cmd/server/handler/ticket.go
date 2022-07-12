package handler

import (
	"net/http"

	"desafio-go-web/internal/tickets"

	"github.com/gin-gonic/gin"
)

type Service struct {
	service tickets.Service
}

func NewService(s tickets.Service) *Service {
	return &Service{
		service: s,
	}
}

func (s *Service) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tickets, err := s.service.GetAll()
		if err != nil {
			ctx.JSON(400, "Ocurrio un error")
			return
		}
		ctx.JSON(200, tickets)
	}
}

func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")
		if destination == "" {
			c.JSON(http.StatusBadRequest, "Destination must be provided")
		}

		tickets, err := s.service.GetTicketByDestination(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tickets)
	}
}

func (s *Service) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")
		if destination == "" {
			c.JSON(http.StatusBadRequest, "Destination must be provided")
		}

		avg, err := s.service.AverageDestination(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}
