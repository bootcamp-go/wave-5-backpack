package users

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
)

var users []domain.ModelUser
var lastId int

type Repository interface {
	GetAll() ([]domain.ModelUser, error)
	GetById(id int) (domain.ModelUser, error)
	Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool) (domain.ModelUser, error)
	LastId() (int, error)
	Update(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool) (domain.ModelUser, error)
	UpdateApellidoEdad(id int, nombre string, edad int) (domain.ModelUser, error)
	Delete(id int) error
}

type repository struct{}

func loadUsers() {
	// Abrimos el archivo
	jsonFile, err := os.Open("./usuarios.json")
	if err != nil {
		users = []domain.ModelUser{}
	}

	// Cerramos el archivo
	defer jsonFile.Close()

	// Creamos el objeto decodificador
	decoder := json.NewDecoder(jsonFile)

	// Decodificamos el archivo y se asigna al slice usuarios
	err = decoder.Decode(&users)
	if err == io.EOF { // El archivo JSON esta vacio
		users = []domain.ModelUser{}
	}
	if err != nil { // No se puede parsear el archivo JSON
		users = []domain.ModelUser{}
	}
}

func NewRepository() Repository {
	loadUsers()
	return &repository{}
}

func (r *repository) GetAll() ([]domain.ModelUser, error) {
	return users, nil
}

func (r *repository) GetById(id int) (domain.ModelUser, error) {
	var user domain.ModelUser
	found := false
	for i := range users {
		if users[i].Id == id && !found {
			user = users[i]
			found = true
		}
	}

	if !found {
		return domain.ModelUser{}, fmt.Errorf("usuario %d no encontrado", id)
	}
	return user, nil
}

func (r *repository) LastId() (int, error) {
	if len(users) > 0 {
		lastId = 0
		for _, u := range users {
			if u.Id > lastId {
				lastId = u.Id
			}
		}
		return lastId, nil
	} else {
		return 0, nil
	}
}

func (r *repository) Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool) (domain.ModelUser, error) {
	u := domain.ModelUser{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaCreacion: time.Now()}
	users = append(users, u)
	return u, nil
}

func (r *repository) Update(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool) (domain.ModelUser, error) {
	user := domain.ModelUser{Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo}
	updated := false
	for i := range users {
		if users[i].Id == id && !updated {
			user.Id = id
			user.FechaCreacion = users[i].FechaCreacion
			users[i] = user
			updated = true
		}
	}

	// Verificamos que se haya actualizado el registro
	if !updated {
		return domain.ModelUser{}, fmt.Errorf("usuario %d no encontrado", id)
	}
	return user, nil
}

func (r *repository) UpdateApellidoEdad(id int, apellido string, edad int) (domain.ModelUser, error) {
	var user domain.ModelUser
	updated := false
	for i := range users {
		if users[i].Id == id && !updated {
			users[i].Apellido = apellido
			users[i].Edad = edad
			user = users[i]
			updated = true
		}
	}

	if !updated {
		return domain.ModelUser{}, fmt.Errorf("usuario %d no encontrado", id)
	}
	return user, nil
}

func (r *repository) Delete(id int) error {
	found := false
	var index int
	for i := range users {
		if users[i].Id == id && !found {
			index = i
			found = true
		}
	}

	if !found {
		return fmt.Errorf("usuario %d no econtrado", id)
	}
	users = append(users[:index], users[index+1:]...)
	return nil
}
