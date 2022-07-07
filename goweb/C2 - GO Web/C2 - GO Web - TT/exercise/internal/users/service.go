package users

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"exercise/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Usuarios, error)
	Store(nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha time.Time) (domain.Usuarios, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Usuarios, error) {
	us, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return us, nil
}

func (s *service) Store(nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha time.Time) (domain.Usuarios, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Usuarios{}, err
	}

	lastID++

	usuario, err := s.repository.Store(lastID, nombre, apellido, email, edad, altura, true, time.Now())
	if err != nil {
		return domain.Usuarios{}, err
	}

	data, err := s.repository.GetAll()
	if err != nil {
		fmt.Println(err)
	}
	jsonData, errMarshal := json.Marshal(data)
	if errMarshal != nil {
		log.Fatal(errMarshal)
	}
	// fmt.Println(data)
	// dataUsers := []byte(string(jsonData))
	os.WriteFile("../users.json", jsonData, 0644)

	return usuario, nil
}
