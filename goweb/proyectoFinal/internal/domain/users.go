package domain

import "time"

type User struct {
	Id       int       `json:"id" binding:"-"`
	Name     string    `json:"name" binding:"required"`
	LastName string    `json:"last_name" binding:"required"`
	Email    string    `json:"email" binding:"required"`
	Age      int       `json:"age" binding:"required"`
	Height   float64   `json:"height" binding:"required"`
	Active   bool      `json:"active" binding:"required"`
	Date     time.Time `json:"creation_date" binding:"required"`
}
