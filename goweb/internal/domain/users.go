package domain

type User struct {
	Id int 				`json:"id"`
	Name string			`json:"name"`
	LastName string		`json:"lastname"`			
	Email string		`json:"email"`
	Age int				`json:"age"`
	Height float32		`json:"height"`
	Active bool			`json:"active"`
	CreatedAt string	`json:"createdat"`
}