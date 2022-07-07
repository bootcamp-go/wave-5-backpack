package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func main() {
	type usuarios struct {
		Id       int       `json:"id" binding:"required"`
		Nombre   string    `json:"nombre"`
		Apellido string    `json:"apellido"`
		Email    string    `json:"email"`
		Edad     int       `json:"edad"`
		Altura   float64   `json:"altura"`
		Activo   bool      `json:"activo"`
		Fecha    time.Time `json:"fecha"`
	}

	//Guardando esta info en users.json
	users := []usuarios{
		{
			Id:       1,
			Nombre:   "Luz",
			Apellido: "Lucumi",
			Email:    "luz.lucumi@hotmail.com",
			Edad:     26,
			Altura:   1.65,
			Activo:   true,
			Fecha:    time.Now(),
		},
		{
			Id:       2,
			Nombre:   "Luber",
			Apellido: "Lucumi",
			Email:    "luber.lucumi@hotmail.com",
			Edad:     61,
			Altura:   1.82,
			Activo:   true,
			Fecha:    time.Now(),
		},
		{
			Id:       3,
			Nombre:   "Martha",
			Apellido: "Hernández",
			Email:    "martha@hotmail.com",
			Edad:     60,
			Altura:   1.60,
			Activo:   true,
			Fecha:    time.Now(),
		},
		{
			Id:       4,
			Nombre:   "Luz",
			Apellido: "Martinez",
			Email:    "luz.martinez@hotmail.com",
			Edad:     26,
			Altura:   1.60,
			Activo:   false,
			Fecha:    time.Now(),
		},
	}

	jsonData, errMarshal := json.Marshal(users)
	if errMarshal != nil {
		log.Fatal(errMarshal)
	}

	dataUsers := []byte(string(jsonData))
	os.WriteFile("../../users.json", dataUsers, 0644)
}
