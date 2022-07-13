package domain

type Ticket struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Country string  `json:"country"`
	Time    string  `json:"time"`
	Price   float64 `json:"price"`
}
