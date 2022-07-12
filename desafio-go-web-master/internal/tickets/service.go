package tickets

type Service interface {
	GetTotalTickets(destination string) (int, error)
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

func (s *service) GetTotalTickets(destination string) (int, error) {
	resp, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return 0, err
	}
	return len(resp), nil
}

func (s *service) AverageDestination(destination string) (float64, error) {
	tickByCountry, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return 0, err
	}
	totalTicket, err := s.repository.GetAll()
	if err != nil {
		return 0, err
	}
	return float64(len(tickByCountry)) / float64(len(totalTicket)), nil
}
