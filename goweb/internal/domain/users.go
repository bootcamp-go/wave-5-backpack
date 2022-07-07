package domain

type Users struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	Age          int     `json:"age"`
	Height       float64 `json:"height"`
	Active       bool    `json:"active"`
	CreationDate string  `json:"creation_date"`
}
