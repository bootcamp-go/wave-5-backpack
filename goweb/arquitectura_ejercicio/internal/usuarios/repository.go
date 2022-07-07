package usuarios

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/anesquivel/wave-5-backpack/goweb/arquitectura_ejercicio/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Usuario, error)
	Store(id, age int, names, lastname, email, dateCreated string, estatura float64) (domain.Usuario, error)
	LastID() (int, error)
	GetUsersFromFile() (domain.Usuarios, error)
	SetFromFile() (bool, error)
}

type repository struct{}

var usersList []domain.Usuario
var lastID int

func (r *repository) GetUsersFromFile() (domain.Usuarios, error) {
	var usersArr domain.Usuarios
	jsonFile, errOpenFile := os.Open("../usuarios.json")
	if errOpenFile != nil {
		fmt.Println("open file", errOpenFile)
		return domain.Usuarios{}, errOpenFile
	}
	fmt.Println("····· Successfully Opened users.json ✅")
	defer jsonFile.Close()

	byteValue, eReadingJsonFile := ioutil.ReadAll(jsonFile)

	if eReadingJsonFile != nil {
		return domain.Usuarios{}, eReadingJsonFile
	}

	eUnmarshal := json.Unmarshal(byteValue, &usersArr)

	if eUnmarshal != nil {
		return domain.Usuarios{}, eUnmarshal
	}

	return usersArr, nil
}
func (r *repository) SetFromFile() (bool, error) {
	usersFromFile, err := r.GetUsersFromFile()

	if err != nil {
		return false, errors.New("Ocurrió un error al intentar setear la lista.")
	}
	fmt.Println("entrara aca")

	usersList = usersFromFile.Users
	return true, nil
}
func (r *repository) GetAll() ([]domain.Usuario, error) {
	if len(usersList) == 0 {
		return []domain.Usuario{}, errors.New("No hay usuarios registrados.")
	}
	return usersList, nil
}

func (r *repository) Store(id, age int, names, lastname, email, dateCreated string, estatura float64) (domain.Usuario, error) {
	nwUsuario := domain.Usuario{
		Id:          id,
		Names:       names,
		LastName:    lastname,
		Age:         age,
		DateCreated: dateCreated,
		Estatura:    estatura,
		Email:       email,
		IsActivo:    true,
	}
	usersList = append(usersList, nwUsuario)
	return nwUsuario, nil
}
func (r *repository) LastID() (int, error) {
	return lastID, nil
}
func NewRepository() Repository {
	return &repository{}
}
