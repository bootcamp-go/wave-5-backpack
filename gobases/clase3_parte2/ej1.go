/* EJ1
Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura. 
Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones:

La estructura debe tener los campos: 
Nombre, 
Apellido, 
edad, 
correo y 
contraseña

Y deben implementarse las funciones:
cambiar nombre: me permite cambiar el nombre y apellido.
cambiar edad: me permite cambiar la edad.
cambiar correo: me permite cambiar el correo.
cambiar contraseña: me permite cambiar la contraseña.

*/

package main

import (
	"fmt"
)

type Usuario struct {
	nombre	string
	apellido	string
	edad	int
	correo	string
	contrasena string	
}



func main() {


	p1 := New("Stefano", "Trejo", 30, "stefanots@gmail.com", "123456")

	//d1 := []byte(string(fmt.Sprint(p1.nombre, "," , p1.apellido)))	

	cambiarNombre(&p1.nombre, &p1.apellido)
	cambiarEdad(&p1.edad)
	cambiarCorreo(&p1.correo)
	cambiarContrasena(&p1.contrasena)

	fmt.Println(p1.nombre, p1.apellido, p1.edad, p1.correo, p1.contrasena)
}


func New(nombreU string, apellidoU string, edadU int, correoU string, contrasenaU string ) Usuario {
	 return Usuario{nombre: nombreU, apellido: apellidoU, edad: edadU, correo: correoU, contrasena: contrasenaU}
}

func cambiarNombre(nombre *string, apellido *string) {
*nombre = "belen"
*apellido = "juarez"
}

func cambiarEdad(edad *int) {
	*edad = 29
	}

	func cambiarCorreo(correo *string) {
		*correo = "mabelennjq@gmail.com"
		}

		func cambiarContrasena(contrasena *string) {
			*contrasena = "45678"
			}
	