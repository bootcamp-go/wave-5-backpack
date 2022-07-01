package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type cliente struct {
	legajo    int
	nombre    string
	apellido  string
	dni       int
	telefono  int
	domicilio string
}

func generarLegajo() int {
	rand.Seed(time.Now().UnixNano())
	legajo := rand.Intn(1000000)
	if legajo < 100000 || legajo > 600000 {
		panic("Error: El legajo debe estar entre 100000 y 600000")
	} else {
		rutaArchivo := "./customers.txt"
		leerArchivo(rutaArchivo)
		return legajo
	}
}

func leerArchivo(rutaArchivo string) string {
	data, err := os.ReadFile(rutaArchivo)

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado.\n")
	} else {
		return string(data)
	}

}

func (c *cliente) validacionDatos(nombre, apellido string, dni, telefono int, domicilio string) (string, error) {
	if nombre != "" && apellido != "" && dni > 0 && telefono != 0 && domicilio != "" {
		c.nombre = nombre
		c.apellido = apellido
		c.dni = dni
		c.telefono = telefono
		c.domicilio = domicilio
		return "Se registró el cliente satisfactoriamente!", nil
	} else {
		return "", errors.New("Error: ingresó algún dato inválido.")
	}
}

func main() {

	defer func() {
		fmt.Println("Fin de la ejecución")
		fmt.Println("Se detectan varios errores en tiempo de ejecución")
		fmt.Println("No han quedado archivos abiertos")
	}()

	legajo := generarLegajo()
	c1 := cliente{legajo: legajo}
	data, err := c1.validacionDatos("Jessica", "Escobar", 123, 316, "Carrera 13")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}

}

// COMENTARIOS

// validacionDatos: Devolviendo una estructura funciona pero no es tan optimizado ya que en caso error se debe devolver una estructura vacia
// y eso ocupa un lugar en memoria innecesario

/* func validacionDatos(legajo int, nombre, apellido string, dni, telefono int, domicilio string) (cliente, error) {
	if nombre != "" && apellido != "" && dni > 0 && telefono != 0 && domicilio != "" {
		c1 := cliente{nombre: nombre, apellido: apellido, dni: dni, telefono: telefono, domicilio: domicilio}
		return c1, nil
	} else {
		return cliente{}, errors.New("Error: ingresó algún dato inválido.")
	}
} */
