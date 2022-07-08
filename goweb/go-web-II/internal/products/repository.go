package products

import (
	"encoding/json"
	"errors"
	"fmt"
	"goweb/go-web-II/internal/domain"
	"io/ioutil"
	"os"
)

/*
En repository va a ir todo lo que sea manejo de BD (Consultas)
y variables globales.
*/




type Repository interface {
	GetAll() ([]*domain.User, error)
	LastId() (int, error)
	Store(id, age int, name, surname, email, created string, active bool) (*domain.User, error)
	Update(id, age int, name, surname, email, created string, active bool )(*domain.User, error)
	Delete(id int) error 
}
type repository struct {}

var users []*domain.User
var lastId int



func NewRepository() Repository {
	return &repository{}
}


func (r *repository)Delete(id int) error {
	deleted := false
	var index int 
	for i := range users {
		if users[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("producto %d no encotrado", id)
	}
	users = append(users[:index], users[index+1:]...)  /*agarra users desde el principio hasta el indice del user encontrado y lo "recorta" y lo sigue desde el indice + 1*/
	return nil
}


func (r *repository)Update(id, age int, name, surname, email, created string, active bool )(*domain.User, error){
	u := &domain.User{Name: name, Surname: surname, Email: email, Created: created, Age: age, Active: active}
	updated := false
	for i := range users {
		if users[i].Id == id {
			u.Id = id
			users[i] = u
			updated = true
		}
	}
	if !updated {
		return &domain.User{}, fmt.Errorf("producto %d no encontrado", id)
	}
	return u, nil
}




func (r *repository) Store(id, age int, name, surname, email, created string, active bool)(*domain.User, error){
	user := &domain.User{
		Id : id,
		Name: name,
		Age: age,
		Surname: surname,
		Email: email,
		Created: created,
		Active: active,
	}
	users = append(users, user)
	lastId = user.Id
	return user, nil
}

func (r *repository) LastId()(int,error){
	return lastId, nil
}

func read()([]*domain.User, error){
	jsonFile, _ := os.Open("usuarios.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var u []*domain.User
	err := json.Unmarshal(byteValue, &u)
	if err != nil {
		return []*domain.User{}, errors.New("tuvimos un problema en la serialización")
	}
	return u, nil
}

func (r *repository) GetAll() ([]*domain.User, error){
	users, _ = read()
	return users, nil
}
