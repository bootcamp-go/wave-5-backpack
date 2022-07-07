package user

type CreateUserRequest struct {
	Id        int     `json:"id"`
	FirstName string  `json:"firstName"  binding:"required"`
	LastName  string  `json:"lastName"  binding:"required"`
	Email     string  `json:"email"  binding:"required"`
	Age       int     `json:"age"  binding:"required"`
	Height    float64 `json:"height"  binding:"required"`
}
