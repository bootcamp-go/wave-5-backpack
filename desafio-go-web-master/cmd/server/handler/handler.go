package handler

import (
	"desafio-go-web-master/internal/tickets"
	"desafio-go-web-master/pkg/web"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ticket struct {
	service tickets.Service
}

func NewProduct(s tickets.Service) *Ticket {
	return &Ticket{
		service: s,
	}
}

func (t *Ticket) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		p, err := t.service.GetAll()
		if err != nil {
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		c.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (t *Ticket) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")
		tickets, err := t.service.GetTicketByDestination(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, web.NewResponse(200, tickets, ""))
	}
}

func (t *Ticket) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := t.service.AverageDestination(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}
		c.JSON(200, web.NewResponse(200, avg, ""))
	}
}
