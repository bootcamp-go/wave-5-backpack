package tickets

import (
	"desafio-go-web/internal/domain"
	"strconv"
)


type Service interface {
	GetAll() ([]domain.Ticket, error)
	GetTotalTickets(destination string)  (string, error)
	GetAverageDestination(destination string) (string, error)

}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Ticket, error) {

	return s.repository.GetAll()
	
}

func (s *service) GetTotalTickets(destination string)  (string, error){

	tickets, err := s.repository.GetTicketByDestination(destination)

	if err != nil {
		return "", err
	}



	return "La cantidad de tickets para "+ destination + " es de " + strconv.Itoa(len(tickets)), nil
	
}

func (s *service) GetAverageDestination(destination string) (string, error){

	tickets, err := s.repository.GetTicketByDestination(destination)
	allTickets, err1 := s.repository.GetAll()

	if err != nil {
		return "", err
	}
	if err1 != nil {
		return "", err1
	}

	cantTotal := len(allTickets)

	avera := (len(tickets)*100)/cantTotal


	return "La cantidad de tickets promedio por día para "+ destination + " es de " + strconv.Itoa(avera)+ "%", nil
}