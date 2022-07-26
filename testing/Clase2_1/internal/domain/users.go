package domain

type User struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	LastName   string  `json:"last_name"`
	Mail       string  `json:"mail"`
	Years      int     `json:"years"`
	Tall       float64 `json:"tall"`
	Enable     bool    `json:"enable"`
	CreateDate string  `json:"create_date"`
}
