package main

import "fmt"

func main() {
	/*//1
	usuario := newUsuario("Francisco", "Monay", 31, "francisco.monay@mercadolibre.com", "123456")

	usuario.detalleU()

	usuario.cambiarNombre("Pablo", "Funes")
	usuario.cambiarEdad(20)
	usuario.cambiarCorreo("pablofunes@gmail.com")
	usuario.cambiarPassword("654321")

	usuario.detalleU()
	*/

	/*//2
	producto := newProducto("Televisor", 100)

	usuario := newUsuario("Francisco", "Monay", 31, "francisco.monay@mercadolibre.com", "123456")

	addProducto(&usuario, &producto, 10)

	fmt.Println(usuario)

	deleteProdUsu(&usuario)

	fmt.Println(usuario)
	*/
	var productos []producto
	var servicios []servicio
	var mantenimientos []mantenimiento

	for i := 0; i < 3; i++ {
		nombreP := fmt.Sprintf("Televisor%d", i)
		producto := newProducto(nombreP, 10, 2)
		productos = append(productos, producto)
	}

	for i := 0; i < 3; i++ {
		nombreS := fmt.Sprintf("Limpieza%d", i)
		servicio := newServicio(nombreS, 10, 30)
		servicios = append(servicios, servicio)
	}

	for i := 0; i < 3; i++ {
		nombreM := fmt.Sprintf("Estructura%d", i)
		mantenimiento := newMantenimiento(nombreM, 10)
		mantenimientos = append(mantenimientos, mantenimiento)
	}

	cP := make(chan float64)
	go sumarProductos(productos, cP)
	cS := make(chan float64)
	go sumarServicios(servicios, cS)
	cM := make(chan float64)
	go sumarMantenimiento(mantenimientos, cM)

	fmt.Println("Monto final: ", (<-cP + <-cS + <-cM))
}
