package domain

type User struct {
	Id, Age                               int
	FirstName, LastName, Email, CreatedAt string
	Height                                float64
	Active                                bool
}
