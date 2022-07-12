package tickets

type Service interface {
	GetCountTicketsByDestination(destination string) (int, error)
	GetPercentageByDestination(destination string) (float64, error)
}

type service struct {
	rep Repository
}

func NewService(rep Repository) Service {
	return &service{
		rep: rep,
	}
}

func (ser *service) GetCountTicketsByDestination(destination string) (int, error) {
	tickets, err := ser.rep.GetTicketByDestination(destination)
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}
func (ser *service) GetPercentageByDestination(destination string) (float64, error) {
	ticketsAll, err := ser.rep.GetAll()
	if err != nil {
		return 0, err
	}
	tickets, err := ser.rep.GetTicketByDestination(destination)
	if err != nil {
		return 0, err
	}
	return float64(len(tickets)) / float64(len(ticketsAll)) * 100, nil
}
