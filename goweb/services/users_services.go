package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

type User struct {
	Id            string  `json:"id"`
	Nombre        string  `json:"nombre"`
	Apellido      string  `json:"apellido"`
	Email         string  `json:"email"`
	Edad          int     `json:"edad"`
	Altura        float64 `json:"altura"`
	Activo        bool    `json:"activo"`
	FechaCreacion string  `json:"fechaCreacion"`
}

func Read() ([]User, error) {
	file, err := ioutil.ReadFile("./users.json")
	if err != nil {
		log.Println(err)
		return nil, errors.New("error al leer el archivo users.json")
	}

	users := []User{}
	if err := json.Unmarshal([]byte(file), &users); err != nil {
		log.Println(err)
		return nil, errors.New("error al deserializar el archivo users.json")
	}

	return users, nil
}
