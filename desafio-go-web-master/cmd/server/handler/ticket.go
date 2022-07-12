package handler

import (
	"desafio-go-web/internal/tickets"
	"desafio-go-web/pkg/web"

	"github.com/gin-gonic/gin"
)

type request struct{
	Name    string
	Email   string
	Country string
	Time    string
	Price   float64
}

type Ticket struct{
	service tickets.Service
}

func  NewTicket(s tickets.Service) *Ticket  {

	return &Ticket{service: s}
	
}

func (t *Ticket) GetAll() gin.HandlerFunc  {

	return func(ctx *gin.Context) {
		t, err := t.service.GetAll()

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		}

		ctx.JSON(200, t)
	}
	
}

func (t *Ticket) GetByCountry() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		req := ctx.Param("dest")

		if req == "" {
			ctx.JSON(404, web.NewResponse(404, nil, "Campo vacio"))
			return
		}

		ts, err := t.service.GetTotalTickets(req)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, ts, ""))
	}
	
}

func (t *Ticket) GetByAverage() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		req := ctx.Param("dest")

		if req == "" {
			ctx.JSON(404, web.NewResponse(404, nil, "Campo vacio"))
			return
		}

		ts, err := t.service.GetAverageDestination(req)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, ts, ""))
	}
	
}

/*import (
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

func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTotalTickets(destination)
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

		avg, err := s.service.AverageDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}*/
