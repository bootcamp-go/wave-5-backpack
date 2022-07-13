package handler

import (
	"log"
	"net/http"
	"os"

	"desafio-go-web/internal/domain"
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
// @Summary Get tickets by country
// @Tags Tickets
// @Description Get tickets by country
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param dest path string true "destiny"
// @Success 200 {object} web.Response
// @Router /ticket/getByCountry/{dest} [get]
func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		destination := c.Param("dest")
		log.Println("destination: ", destination)
		if destination == "" {
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, []domain.Ticket{}, nil))
		}

		tickets, err := s.service.GetTotalTickets(c, destination)
		if err != nil {
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, err.Error(), nil))
			return
		}

		c.JSON(200, tickets)
	}
}

// AverageDestination godoc
// @Summary Get average by country
// @Tags Tickets
// @Description Get average by country
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param dest path string true "destiny"
// @Success 200 {object} web.Response
// @Router /ticket/getAverage/{dest} [get]
func (s *Service) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.AverageDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, avg, nil))
	}
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("no esta definido el token de seguridad")
	}

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token required"))
			return
		}

		if token != requiredToken {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "wrong token"))
			return
		}

		ctx.Next()
	}
}
