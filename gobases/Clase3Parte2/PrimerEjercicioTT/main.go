package main

import "fmt"

type Persona struct {
	Nombre      string
	Apellido    string
	Edad        int
	Correo      string
	Contrasenia string
}

func main() {
	p := Persona{"Jose", "Gonzalez", 22, "jvgp@meli.com", "contraseña"}
	fmt.Print(p)
	cambiarNombre("Vicente", "Perez", &p)
	cambiarEdad(23, &p)
	cambiarCorreo("pgvj@lime.moc", &p)
	cambiarContrasenia("contrasenia", &p)
	fmt.Print(p)
}

func cambiarNombre(nombre string, apellido string, persona *Persona) {
	persona.Nombre = nombre
	persona.Apellido = apellido
}
func cambiarEdad(edad int, persona *Persona) {
	persona.Edad = edad
}
func cambiarCorreo(correo string, persona *Persona) {
	persona.Correo = correo
}
func cambiarContrasenia(contr string, persona *Persona) {
	persona.Contrasenia = contr
}
