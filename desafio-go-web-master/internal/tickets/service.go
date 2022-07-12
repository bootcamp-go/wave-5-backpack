package tickets

import (
"fmt"
)

type Service interface{
	GetTotalTickets()
	GetTicketsByCountry()
	AverageDestination()
}

type service struct{
	repository Repository
} 

func NewService (r Repository) Service{
	return &service{
		repository: r,
	}
}
func (s *service) GetAll() ([]domain .Ticket, error) {
	tk, err := s.repository.GetAll()
	fmt.Println(tk)
	if err != nil {
		return nil, err
	}
	return tk, nil
}
