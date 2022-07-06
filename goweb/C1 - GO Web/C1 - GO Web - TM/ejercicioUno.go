package main

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func main() {
	type usuarios struct {
		Id                      int `json:"id" binding:"required"`
		Nombre, Apellido, Email string
		Edad                    int
		Altura                  float64
		Activo                  bool
		Fecha                   time.Time
	}

	users := []usuarios{
		{
			Id:       1,
			Nombre:   "Luz",
			Apellido: "Lucumí",
			Email:    "luz.lucumi@hotmail.com",
			Edad:     26,
			Altura:   1.65,
			Activo:   true,
			Fecha:    time.Now(),
		},
		{
			Id:       2,
			Nombre:   "Luber",
			Apellido: "Lucumí",
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
	}

	jsonData, errMarshal := json.Marshal(users)
	if errMarshal != nil {
		log.Fatal(errMarshal)
	}

	dataUsers := []byte(string(jsonData))
	os.WriteFile("./users.json", dataUsers, 0644)
}
