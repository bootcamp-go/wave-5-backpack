package tickets

import (
	"desafio-go-web/internal/domain"
	"desafio-go-web/pkg/store"
	"fmt"
)

var ts []domain.Ticket

type Repository interface {
	GetAll() ([]domain.Ticket, error)
	GetTicketByDestination(destination string) ([]domain.Ticket, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}
func (r *repository) GetAll() ([]domain.Ticket, error) {
	err := r.db.Read(&ts)

	if err != nil {
		return []domain.Ticket{}, fmt.Errorf(err.Error())
	}

	return ts, nil
}

/*func (r *repository) GetAll() ([]domain.Ticket, error) {
	var ticketList []domain.Ticket

	file, err := os.Open("./tickets.csv")
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}*/

func (r *repository) GetTicketByDestination(destination string) ([]domain.Ticket, error) {

	tickets, err := r.GetAll()
	var ticketsDest []domain.Ticket

	if err != nil {
		return []domain.Ticket{}, fmt.Errorf("empty list of tickets")
	}

	for _, t := range tickets {
		if t.Country == destination {
			ticketsDest = append(ticketsDest, t)
		}
	}

	return ticketsDest, nil
}
