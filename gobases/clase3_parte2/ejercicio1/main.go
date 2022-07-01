package main

import "fmt"

func main() {
	//definida dentro del mail la var estructura
	//parametro de funcion un puntero de la estructura.

	type Usuarios struct {
		Nombre      string
		Apellido    string
		Edad        int
		correo      string
		contrasenia string
	}

	u := Usuarios{
		Nombre:      "Yvo",
		Apellido:    "Pintos",
		Edad:        30,
		correo:      "yvonnepintos",
		contrasenia: "12345",
	}
	cambiarNombre(&u.Nombre)
	fmt.Println("El nuevo nombre es:", u.Nombre)
	cambiarEdad(&u.Edad)
	fmt.Println("La nueva edad es:", u.Edad)
	cambiarCorreo(&u.correo)
	fmt.Println("El nuevo correo es:", u.correo)
	cambiarContrasenia(&u.contrasenia)
	fmt.Println("La nueva clave es:", u.contrasenia)
}

func cambiarNombre(nombre *string) {
	*nombre = "fruta"
}

func cambiarEdad(edad *int) {
	*edad = 44
}

func cambiarCorreo(correo *string) {
	*correo = "blabla@bla.com"
}

func cambiarContrasenia(contra *string) {
	*contra = "324234234234"
}
