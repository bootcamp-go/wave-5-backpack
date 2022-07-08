package main


import (
	"fmt"
)


var (
	pequenio string = "PEQUEÑO"
	mediano  string = "MEDIANO"
	grande   string = "GRANDE"
)

//Crear una estructura “tienda” que guarde una lista de productos. 
type tienda struct {
	prod producto
}

// Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
type producto struct {   
	tipo	 string
	nombre 	 string
	precio	 float64
}


type Producto interface {
	CalcularCosto ()
}

func (this producto) CalcularCosto(){

}

type Ecommerce interface{
	Total() 
	Agregar()
}

// El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
func (this tienda) total() float64{
	switch this.prod.tipo {
	case pequenio:
		return this.prod.precio
	case mediano:
		return this.prod.precio * 1.03 
	case grande:
		return (this.prod.precio * 1.06)+2.500 

	default: return 0
	}

}
// El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda
func (this *tienda) agregar(){
	

}

/*
func (u *Usuario) AgregarProducto(producto *Producto, cantidad *int) {
	producto.Cantidad = *cantidad
	u.Productos = append(u.Productos, *producto)
}
*/

func main() {

p1 := producto{tipo: pequenio, nombre: "choclo", precio: 100.00}
p2 := producto{grande, "licuadora", 2000.00}

listaTienda := []producto{p1,p2}

listaTienda = append(listaTienda, )


fmt.Println(p1)
fmt.Println(p2)
}


//Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
func  nuevoProducto(tipoProducto string , nombre string, precio float64) producto{
	return producto{tipoProducto, nombre, precio}

}

// Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
/*
func nuevaTienda (productos... Producto) Ecommerce{

	return Tienda{}
	
}*/