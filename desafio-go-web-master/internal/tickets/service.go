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
	tickets, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}

func (s *service) AverageDestination(destination string) (float64, error) {
	tickets, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return 0, nil
	}
	allTickets, err := s.repository.GetAll()
	if err != nil {
		return 0, nil
	}
	avg := float64(len(tickets)) / float64(len(allTickets))
	return avg, nil
}
