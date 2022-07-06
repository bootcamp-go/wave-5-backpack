package main

import "fmt"

type producto struct {
	nombre   string
	precio   int
	cantidad int
}

func nuevoProducto(nombre string, precio int) producto {
	p := producto{nombre, precio, 0}
	return p
}

func sumar_prod(c chan int, list_prod []producto) {
	precio_total := 0
	for _, prod := range list_prod {
		precio_total += prod.precio * prod.cantidad
	}
	c <- precio_total
}

type servicio struct {
	nombre      string
	precio      int
	minutosTrab int
}

func sumarServicio(c chan int, list_servio []servicio) {
	precioTotal := 0
	for _, servicio := range list_servio {
		mediasHoras := servicio.minutosTrab / 30
		if servicio.minutosTrab%30 != 0 {
			mediasHoras++
		}
		precioTotal += servicio.precio * mediasHoras
	}
	c <- precioTotal
}

type mantenimiento struct {
	nombre string
	precio int
}

func sumar_mant(c chan int, list_mant []mantenimiento) {
	precio_total := 0
	for _, mant := range list_mant {
		precio_total += mant.precio
	}
	c <- precio_total
}

type usuario struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contrasena string
	productos  []*producto
}

func (u *usuario) cambiar_nom(nue_nombre, nue_apellido string) {
	u.nombre = nue_nombre
	u.apellido = nue_apellido
}
func (u *usuario) cambiar_edad(nue_edad int) {
	u.edad = nue_edad
}
func (u *usuario) cambiar_correo(nue_correo string) {
	u.correo = nue_correo
}
func (u *usuario) cambiar_contr(nue_contrasena string) {
	u.contrasena = nue_contrasena
}

func agregaProducto(usur *usuario, prod *producto, cant int) {

	(*prod).cantidad = cant
	(*usur).productos = append((*usur).productos, prod)
}

func borrarProductos(usur *usuario) {
	(*usur).productos = nil
}

func main() {
	// u1 := usuario{"juan", "de la pena", 6, "po@live.cl", "tumama", nil}
	// fmt.Println(u1)
	// u1.cambiar_nom("herm", "t1viej")
	// fmt.Println(u1)
	// u1.cambiar_edad(67)
	// fmt.Println(u1)
	// prod1 := nuevoProducto("silla", 15000)
	// prod2 := nuevoProducto("manzana", 150)
	// fmt.Println(prod1)
	// agregaProducto(&u1, &prod1, 23)
	// agregaProducto(&u1, &prod2, 5)
	// fmt.Println(prod1.cantidad)
	// fmt.Println((*u1.productos[1]).nombre)
	// borrarProductos(&u1)
	// fmt.Println((u1).productos)
	product1 := producto{"bicicleta", 2500, 5}
	product2 := producto{"tv", 900, 2}
	list_prod := []producto{product1, product2}
	c := make(chan int)
	go sumar_prod(c, list_prod)
	serv1 := servicio{"pintar", 3000, 130}
	serv2 := servicio{"tapisar", 5000, 220}
	listServ := []servicio{serv1, serv2}
	go sumarServicio(c, listServ)
	mant1 := mantenimiento{"atornillar", 4000}
	mant2 := mantenimiento{"re-pintar", 5000}
	listMant := []mantenimiento{mant1, mant2}
	go sumar_mant(c, listMant)
	fmt.Println("sum all")
	total := 0
	for i := 0; i < 3; i++ {
		mensaje := <-c
		fmt.Println("llego el numero ", mensaje)
		total += mensaje
		fmt.Println("esta es la suma ", total)
	}
}
