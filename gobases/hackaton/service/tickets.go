package service

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(T []Ticket) Bookings {
	return &bookings{
		Tickets: T,
	}
}

func (b *bookings) Create(newData []Ticket, t []Ticket) (Ticket, error) {
	/* p1 := Ticket{1001, "Juana", "correo@correo.com", "np", "12-10-2022", 3} */
	t = append(t, newData)
	return Ticket{}, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	return Ticket{}, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	return Ticket{}, nil
}

func (b *bookings) Delete(id int) (int, error) {
	return 0, nil
}
