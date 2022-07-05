package main

import "fmt"

type Productos struct{
	nombre string
	precio float64
	cantidad int
}
func NuevoProducto(id string, prec float64,cant int) Productos{
	prod:=Productos{
		nombre: id,
		precio: prec,
		cantidad:cant,
	}
	return prod
}

func SumarMantenimiento(listaM *[]Mantenimiento, c chan float64) {
	var total=0.0
	for _,mant :=range *listaM {
		total=total+mant.precio
	}
	fmt.Printf("termino mantenimiento: %v\n",total)
	c <- total
}

func SumarProductos(listaP *[]Productos, c chan float64) {
	var total=0.0
	for _,prod :=range *listaP {
		total=total+(float64(prod.cantidad)*prod.precio)
	}
	fmt.Printf("termino productos: %v\n",total)
	c <- total
}

func SumarServicio(listaS *[]Servicios, c chan float64){
	var total=0.0
	for _,serv :=range *listaS {
		var mediasHorasTrabajadas int
		mediasHorasTrabajadas=serv.minutosTrabajados/30
		if(float64(mediasHorasTrabajadas)<float64(serv.minutosTrabajados)/30.0){
			mediasHorasTrabajadas++
		}
		total=total+(float64(mediasHorasTrabajadas)*serv.precio)
	}
	fmt.Printf("termino servicio: %v\n",total)
	c <- total
}

type Servicios struct{
	nombre string
	precio float64
	minutosTrabajados int
}

func NuevoServicio(id string, prec float64,minutos int) Servicios{
	serv:=Servicios{
		nombre: id,
		precio: prec,
		minutosTrabajados:minutos,
	}
	return serv
}

type Mantenimiento struct{
	nombre string
	precio float64
}

func NuevoMantenimiento(id string, prec float64) Mantenimiento{
	mant:=Mantenimiento{
		nombre: id,
		precio: prec,
	}
	
	return mant
}

func main(){
	var listaP []Productos
	var listaS []Servicios
	var listaM []Mantenimiento
	prod1:=NuevoProducto("Triciclo",3500,1)
	prod2:=NuevoProducto("Bicicleta",5000,1)
	s1:=NuevoServicio("s1",3500,30)
	s2:=NuevoServicio("s2",5000,45)
	m1:=NuevoMantenimiento("m1",3500)
	m2:=NuevoMantenimiento("m2",5000)
	listaP=append(listaP,prod1)
	listaP=append(listaP,prod2)
	listaS=append(listaS,s1)
	listaS=append(listaS,s2)
	listaM=append(listaM,m1)
	listaM=append(listaM,m2)
	cP:=make(chan float64)
	cM:=make(chan float64)
	cS:=make(chan float64)
	go SumarMantenimiento(&listaM,cM)
	go SumarProductos(&listaP,cP)
	go SumarServicio(&listaS,cS)
	//fmt.Printf("Suma Productos %v\n",<-cP)
	//fmt.Printf("Suma Servicios %v\n",<-cS)
	//fmt.Printf("Suma Mantenimiento %v\n",<-cM)
	sumaM:=<-cM
	sumaP:=<-cP
	sumaS:=<-cS
	fmt.Printf("Suma total=%v\n",sumaM+sumaP+sumaS)
}





/*
Ejercicio 3 - Calcular Precio
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, 
Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que 
el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, 
	si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).

*/