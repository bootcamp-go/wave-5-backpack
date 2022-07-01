package main

import "fmt"

func main(){
	type Usuario struct{
		nombre string
		apellido string
		correo string
		prod []prod
	}

	type prod struct{
		nombre string
		precio int
		cantidad int
	}
	p1:= prod{
		nombre: "mate",
		precio: 234,
		cantidad: 2,
	}
	fmt.Println("El valor de los productos es:", p1.nombre, p1.precio)
nuevoProducto(&p1.nombre, &p1.precio)

}

func nuevoProducto(pprod *prod) p prod{
	*nombre = "Otro Nombre"
	*precio = 34
	return 
}

func agregarProducto()