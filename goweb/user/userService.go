package user

import "time"

func (user *User) createUser(u CreateUserRequest) {
	user.Id = u.Id
	user.LastName = u.LastName
	user.FirstName = u.FirstName
	//user.Age = u.Age
	user.Email = u.Email
	//user.Height = u.Height
	user.CreationDate = time.Now().Format("2006-01-02")
	user.Active = true

}
