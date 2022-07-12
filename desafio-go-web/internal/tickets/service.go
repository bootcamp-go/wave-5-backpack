package tickets

type Service interface {
	GetTicketsByCountry(string) (int, error)
	AverageDestination(string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTicketsByCountry(destination string) (int, error) {
	ts, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return 0, err
	}
	return len(ts), nil
}

func (s *service) AverageDestination(destination string) (float64, error) {
	ts, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return 0, err
	}
	tsTotal, err := s.repository.GetAll()
	if err != nil {
		return 0, err
	}
	return float64(len(ts)) / float64(len(tsTotal)), nil
}
