package handler

import (
	"net/http"

	"desafio-go-web/internal/tickets"
	"desafio-go-web/pkg/web"

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

// GetTicketsByCountry godoc
// @Summary Get a ticket total by destination
// @Tags Tickets
// @Description Get a ticket total by destination
// @Produce  json
// @Param dest path int true "Destination"
// @Success 200 {object} web.Response
// @Error 500 {object} web.Response
// @Router /getByCountry/{dest} [get]
func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetCountTicketsByDestination(destination)
		if err != nil {
			c.JSON(web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}
		c.JSON(web.NewResponse(http.StatusOK, gin.H{"destination": destination, "total": tickets}, ""))
	}
}

// PercentageByDestination godoc
// @Summary Get a ticket percentage by destination
// @Tags Tickets
// @Description Get a ticket percentage by destination
// @Produce  json
// @Param dest path int true "Destination"
// @Success 200 {object} web.Response
// @Error 500 {object} web.Response
// @Router /getPercentage/{dest} [get]
func (s *Service) PercentageByDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		perc, err := s.service.GetPercentageByDestination(destination)
		if err != nil {
			c.JSON(web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}

		c.JSON(web.NewResponse(http.StatusOK, gin.H{
			"percentage":  perc,
			"destination": destination,
		}, ""))
	}
}
