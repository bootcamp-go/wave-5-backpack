package main

import "fmt"

func main() {
	/*//1
	leerArchivo("customers.txt")

	fmt.Println("Se controlo el panic y finalizamos el programa correctamente.")
	*/

	nLegajo, err := generarNumeroLegajo()

	if err != nil {
		panic("Numero de legajo invalido")
	}

	verificarCliente(nLegajo)

	cliente := newCliente(nLegajo, "", "Monay", "12343231", "15432345", "Bolivar 123")

	cliente, err = controlDatos(cliente)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Fin de la ejecucion")
	fmt.Println("Se detectaron varios errores en tiempo de ejecucion")
	fmt.Println("No han quedado archivo abiertos")

}
