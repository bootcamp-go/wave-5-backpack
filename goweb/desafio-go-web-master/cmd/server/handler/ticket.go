package handler

import (
	"desafio-go-web/internal/tickets"
	"desafio-go-web/pkg/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	service tickets.Service
}

func NewTicketHandler(s tickets.Service) *TicketHandler {
	return &TicketHandler{
		service: s,
	}
}

// @GetTicketsByCountry godoc
// @Summary Tickets By Country
// @Tags Tickets
// @Description Get all tickets from the selected country
// @Produce json
// @Param token header string true "token"
// @Param dest path string true "country"
// @Success 200 {object} web.Response
// @Router /ticket/getByCountry/{dest} [get]
func (t *TicketHandler) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		destination := c.Param("dest")
		if destination == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(nil, "missing destination"))
		}

		tickets, err := t.service.GetTicketByDestination(destination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(nil, "internal error"))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(tickets, ""))
	}
}

// @AverageDestination godoc
// @Summary Average Tickets By Country
// @Tags Tickets
// @Description Get average number of tickets from the selected country
// @Produce json
// @Param token header string true "token"
// @Param dest path string true "country"
// @Success 200 {object} web.Response
// @Router /ticket/getAverage/{dest} [get]
func (t *TicketHandler) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {
		destination := c.Param("dest")
		if destination == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(nil, "missing destination"))
		}

		avg, err := t.service.GetAverageByDestination(destination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(nil, "internal error"))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(avg, ""))
	}
}
