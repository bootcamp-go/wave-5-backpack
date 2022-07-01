package main

import (
	"fmt"
	"reflect"
)

type Usuarios struct {
	nombre     string
	edad       int
	correo     string
	contraseña string
}

func main() {

	/* var puntero *Usuarios */
	user1 := Usuarios{"asd", 12, "asdas", "sadw"}

	/* user1 = &puntero */

	fmt.Println(reflect.TypeOf(user1))

}
