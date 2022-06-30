package main

type Usuario struct {
	nombre      string
	apellido    string
	edad        int
	correo      string
	constraseña string
}

func (u *Usuario) cambiarNombre(nombre string) {
	u.nombre = nombre
}
func (u *Usuario) cambiarEdad(edad int) {
	u.edad = edad
}
func (u *Usuario) cambiarCorreo(correo string) {
	u.correo = correo
}
func (u *Usuario) cambiarConstraseña(constraseña string) {
	u.constraseña = constraseña
}

func main() {

}
