package handler

import (
	"fmt"
	"net/http"

	"desafio-go-web/internal/tickets"
	"desafio-go-web/pkg/web"

	"github.com/gin-gonic/gin"
)

type Tickets struct {
	service tickets.Service
}

func NewTickets(t tickets.Service) *Tickets {
	return &Tickets{
		service: t,
	}
}

func (t *Tickets) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		// destination := c.Param("dest")

		// tickets, err := t.service.GetTotalTickets(c, destination)
		err := fmt.Errorf("error")
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}
		c.JSON(200, nil)
	}
}

func (t *Tickets) GetTicketByDestination() gin.HandlerFunc {
	return func(c *gin.Context) {
		destination := c.Param("dest")
		tickets, err := t.service.GetTicketByDestination(c, destination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(401, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, tickets, ""))
	}
}

func (t *Tickets) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		t, err := t.service.GetAll()
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, t, ""))
		// c.JSON(200, t)
	}
}

func (t *Tickets) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {
		destination := c.Param("dest")
		tickets, err := t.service.AverageDestination(c, destination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(401, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, tickets, ""))
	}
}

func (t *Tickets) GetTicketDestinationNum() gin.HandlerFunc {
	return func(c *gin.Context) {
		destination := c.Param("dest")
		tickets, err := t.service.GetTicketDestinationNum(c, destination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(401, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, tickets, ""))
	}
}
