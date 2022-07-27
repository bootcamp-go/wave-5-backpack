package users

import (
	"fmt"
	"goweb/internal/domain"
	"goweb/pkg/store"
)


//var lastID int

const (
	UserNotFound = "user %d not found"
	FailReading     = "cant read database"
	FailWriting     = "cant write database"
)


type Repository interface {
	GetAllUsers() ([]domain.User, error)
	GetUserById(id int) (domain.User, error)
	StoreUser(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error)
	LastID() (int,error)
	UpdateTotal(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error)
	UpdatePartial(id int, lastname string, age int) (domain.User, error)
	Delete(id int) error
}

type repository struct{
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return users, nil
}


func (r *repository) GetUserById(id int) (domain.User, error) {
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return domain.User{}, fmt.Errorf(FailReading)
	}

	var userFounded domain.User
	find := false
	for _,u :=range users{
		if u.Id == id{
			userFounded = u
			find = true
			break
		}
	}
	if !find{
		return domain.User{}, fmt.Errorf("no existe el usuario con id %d", id)
	}
	return userFounded, nil
}

func (r *repository) LastID() (int, error) {
	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(users) == 0 {
		return 0, nil
	}

	return users[len(users)-1].Id, nil
}


func (r *repository) StoreUser(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error) {
	var users []domain.User

	if err := r.db.Read(&users); err != nil {
		return domain.User{}, fmt.Errorf(FailReading)
	}

	user := domain.User{Id: id, Name:name, LastName: lastname, Email: email, Age: age, Height: height, Active: active, CreatedAt: createdat}
	users = append(users, user)

	if err := r.db.Write(users); err != nil {
		return domain.User{}, fmt.Errorf(FailWriting)
	}

	return user, nil
}


func(r *repository) UpdateTotal(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error) {

	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return domain.User{}, fmt.Errorf(FailReading)
	}
	
	userToUpdate := domain.User{Name: name, LastName: lastname, Email: email, Age: age, Height: height, Active: active, CreatedAt: createdat}
	updated := false
	
	for i:= range users {
		if users[i].Id == id{
			userToUpdate.Id = id
			users[i] = userToUpdate
			updated = true
			break
		}
	}
	if !updated{
		return domain.User{}, fmt.Errorf(UserNotFound, id)
	}

	if err := r.db.Write(users); err != nil {
		return domain.User{}, fmt.Errorf(FailWriting)
	}

	return userToUpdate, nil
}

func(r *repository) UpdatePartial(id int, lastname string, age int) (domain.User, error) {

	var users []domain.User
	if err := r.db.Read(&users); err != nil {
		return domain.User{}, fmt.Errorf(FailReading)
	}

	updated := false
	var userUpdated domain.User 
	for i:= range users {
		if users[i].Id == id{
			users[i].LastName = lastname
			users[i].Age = age
			updated = true
			userUpdated = users[i]
			break
		}
	}
	if !updated{
		return domain.User{}, fmt.Errorf(UserNotFound, id)
	}

	if err := r.db.Write(users); err != nil {
		return domain.User{}, fmt.Errorf(FailWriting)
	}

	return userUpdated, nil
}


func (r *repository) Delete(id int) error {
   
    var users []domain.User
    if err := r.db.Read(&users); err != nil {
        return fmt.Errorf(FailReading)
    }

    var indexToDelete int
    find := false
    for i :=range users{
        if users[i].Id == id{
            indexToDelete = i
            find = true
            break
        }
    }
    if !find{
        return fmt.Errorf(UserNotFound, id)
    }
    users = append(users[:indexToDelete],users[indexToDelete+1:]... )

    if err := r.db.Write(users); err != nil {
        return fmt.Errorf(FailWriting)
    }
    
    return nil
}