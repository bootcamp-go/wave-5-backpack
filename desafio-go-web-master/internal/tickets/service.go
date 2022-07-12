package tickets

import (
	"desafio-go-web/internal/domain"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetTotalTickets(c *gin.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(c *gin.Context, destination string) (string, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTotalTickets(c *gin.Context, destination string) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetTicketByDestination(c, destination)
	if err != nil {
		return nil, fmt.Errorf("no se pudo obtener los tickets por destino: %w", err)
	}
	return tickets, nil
}

func (s *service) AverageDestination(c *gin.Context, destination string) (string, error) {
	tickets, err := s.repository.GetAll(c)
	if err != nil {
		return "", fmt.Errorf("no se pudo obtener los tickets: %w", err)
	}

	if len(tickets) == 0 {
		return "", fmt.Errorf("no se encontraron tickets")
	}

	ticketsDestino, err := s.repository.GetTicketByDestination(c, destination)
	if err != nil {
		return "", fmt.Errorf("no se pudo obtener los tickets por destino: %w", err)
	}

	// Asumiré que el promedio de personas será el total de tickets de un destino
	// dividido entre el total de tickets de todos los destinos
	// Pero esto me dará un decimal, por lo que lo multiplico por 100 para obtener el porcentaje

	sumX := float64(len(ticketsDestino))
	total := float64(len(tickets))
	avg := (sumX / total) * 100.0

	return fmt.Sprintf("%.1f%%", avg), nil

}
