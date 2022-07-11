package domain

type User struct {
	ID         int     `json: "id"`
	Name       string  `json: "name"`
	Lastname   string  `json: "lastname"`
	Email      string  `json: "email"`
	Age        int     `json: "age"`
	Height     float32 `json: "height"`
	Active     bool    `json: "active"`
	DoCreation string  `json: "doCreation"`
}