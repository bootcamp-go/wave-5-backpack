package user

type User struct {
	Id                         int
	Age                        int    `binding:"required"`
	FirstName, LastName, Email string `binding:"required"`
	CreatedAt                  string
	Height                     float64 `binding:"required"`
	Active                     bool    `binding:"required"`
}
