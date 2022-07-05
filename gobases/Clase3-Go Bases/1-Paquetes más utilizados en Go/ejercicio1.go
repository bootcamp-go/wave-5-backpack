package main

import "fmt"
import "os"

func nuevoProducto(id string, precio float64, cantidad int) producto{
	prod:=producto{
		Id: id,
		Precio: precio,
		Cantidad: cantidad,
	}
	return prod
}


type producto struct{
	Id string
	Precio float64
	Cantidad int
}

func main(){
	prod1:=nuevoProducto("a",34.5,5)
	prod2:=nuevoProducto("b",35.3,24)
	prod3:=nuevoProducto("c",5.3,200)

	var listaProductos []producto
	listaProductos=append(listaProductos,prod1)
	listaProductos=append(listaProductos,prod2)
	listaProductos=append(listaProductos,prod3)

	fmt.Println(listaProductos)
	var salida string
	for _,prod :=range listaProductos {
		salida=salida+fmt.Sprint(prod.Id,",",prod.Precio,",",prod.Cantidad,"\n")
	}
	fmt.Print(salida)

	archivo:=[]byte(salida)
	err:=os.WriteFile("./salida.csv",archivo,0644)
	fmt.Println(err)
}

/*
Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita: 
Implementar una funcionalidad para guardar un archivo de texto, con la 
información de productos comprados, separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.

*/