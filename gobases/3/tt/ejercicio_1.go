package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Usuario struct {
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Edad       int    `json:"edad"`
	Correo     string `json:"correo"`
	Contraseña string `json:"contraseña"`
}

func main() {
	u := &Usuario{
		Nombre:     "Agustin",
		Apellido:   "Nigrelli",
		Edad:       32,
		Correo:     "agustin.nigrelli@mercadolibre.com",
		Contraseña: "123abc",
	}

	antes, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(antes))

	nu := &Usuario{
		Nombre:     "Francisco",
		Apellido:   "Fernandez",
		Edad:       30,
		Correo:     "francisco.fernandez@mercadolibre.com",
		Contraseña: "456def",
	}

	CambiarNombre(&u.Nombre, &u.Apellido, nu.Nombre, nu.Apellido)
	CambiarEdad(&u.Edad, nu.Edad)
	CambiarCorreo(&u.Correo, nu.Correo)
	CambiarContraseña(&u.Contraseña, nu.Contraseña)

	despues, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(despues))
}

func CambiarNombre(nombre *string, apellido *string, nuevo_nombre string, nuevo_apellido string) {
	*nombre = nuevo_nombre
	*apellido = nuevo_apellido
}

func CambiarEdad(edad *int, nueva_edad int) {
	*edad = nueva_edad
}

func CambiarCorreo(correo *string, nuevo_correo string) {
	*correo = nuevo_correo
}

func CambiarContraseña(contraseña *string, nueva_contraseña string) {
	*contraseña = nueva_contraseña
}
