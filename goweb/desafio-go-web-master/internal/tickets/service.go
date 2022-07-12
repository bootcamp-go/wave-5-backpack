package tickets

type Service interface {
	GetTicketByCountry(destination string) (int, error)
	AverageDestination(destination string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTicketByCountry(destination string) (int, error) {

	resp, err := s.repository.GetTicketByCountry(destination)
	if err != nil {
		return 0, err
	}

	return len(resp), nil
}

func (s *service) AverageDestination(destination string) (float64, error) {

	ticketsByCountry, err := s.repository.GetTicketByCountry(destination)

	if err != nil {
		return 0, err
	}

	totalTickets, err := s.repository.GetAll()
	if err != nil {
		return 0, err
	}

	return float64(len(ticketsByCountry)) / float64(len(totalTickets)), nil

}
