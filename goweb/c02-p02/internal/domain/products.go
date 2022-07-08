package domain

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}
