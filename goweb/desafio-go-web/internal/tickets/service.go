package tickets

type Service interface {
	GetTotalTickets(destination string) (int, error)
	AverageDestination(destination string) (float64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetTotalTickets(destination string) (int, error) {
	resp, err := s.repo.GetTicketByDestination(destination)
	if err != nil {
		return 0, err
	}
	return len(resp), nil
}

func (s *service) AverageDestination(destination string) (float64, error) {
	ticketCountry, err := s.repo.GetTicketByDestination(destination)
	if err != nil {
		return 0, err
	}
	totalTicket, err := s.repo.GetAll()
	if err != nil {
		return 0, nil
	}
	return float64(len(ticketCountry)) / float64(len(totalTicket)), nil
}
