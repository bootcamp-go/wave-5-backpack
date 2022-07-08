/*
Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita: 
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
Debe tener:
 el id del producto, 
 precio y la 
 cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.


*/

/*
package main


import (
	"fmt"
	"os"
)


func main() {


	p1 := New(1,100, 2)

	//d1 := []byte(string(p1.id + p1.precio + p1.cantidad))	
	//d1 := []byte(string(fmt.Sprint(p1.id, ",", p1.precio , "," p1.cantidad "\n")))	
	d1 := []byte(string(fmt.Sprint(p1.id, "," , p1.precio , "," ,  p1.cantidad , "\n")))	

	err := os.WriteFile("./myFile.csv", d1, 0644)

	fmt.Println(err)
	if err == nil{
		//fmt.Println(p1.precio)
		data, _ := os.ReadFile("./myFile.csv")
		fmt.Println(string(data))
	
	}
}


type Producto struct {
	id       int
	precio float64
	cantidad int
}

func New(idProducto int, precioProducto float64, cantidadProducto int) Producto {
	 return Producto{id: idProducto, precio: precioProducto, cantidad: cantidadProducto}
}

*/