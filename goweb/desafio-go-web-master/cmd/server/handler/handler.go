package handler

import (
	"desafio-go-web-master/internal/tickets"
	"desafio-go-web-master/pkg/store/web"
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
			//c.JSON(404, gin.H{"error": err.Error()})
			c.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		//c.JSON(200, p)
		c.JSON(200, web.NewResponse(200, p, ""))
	}
}

// func (t *Ticket) GetTotalTickets() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 	}
//}

func (t *Ticket) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")
		tickets, err := t.service.GetTicketByDestination(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, web.NewResponse(200, tickets, ""))
		//c.JSON(200, tickets)
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

		//c.JSON(200, avg)
		c.JSON(200, web.NewResponse(200, avg, ""))
	}
}

//GetAll() ([]domain.Ticket, error)
//GetTotalTickets(destination string) ([]domain.Ticket, error)
