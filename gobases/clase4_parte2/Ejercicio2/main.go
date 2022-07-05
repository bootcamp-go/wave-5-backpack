package main

import (
	"errors"
	"fmt"
	"os"
)

type Cliente struct {
	Nombre, Apellido, Domicilio string
	Legajo, DNI, Telefono       int64
}

func generarLegajo() (int64, error) {
	// retorna nil para fallar
	return 0, errors.New("No se pudo generar el legajo")
}

func leerArchivoClientes() (string, error) {
	file, err := os.ReadFile("customers.txt")

	defer func() {
		errPanic := recover()

		if errPanic != nil {
			fmt.Errorf("error: $s\n", errPanic)
		}
	}()

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	} else {
		return string(file), nil
	}
}

func verificarExistencia() (bool, error) {
	_, err := leerArchivoClientes()
	if err != nil {
		return false, fmt.Errorf("hubo un error al verificar la existencia: %s\n", err)
	}
	//faltaría chequear si existe en la lista o no:
	return false, nil
}

func valoresValidos(nombre, apellido, domicilio string, dni, telefono int64) (bool, error) {
	condicion := nombre == "" || apellido == "" || domicilio == "" || dni == 0 || telefono == 0

	if condicion {
		return false, fmt.Errorf("alguno de los valores ingresados no es válido (%s,%s,%s,%d,%d)\n", nombre, apellido, domicilio, dni, telefono)
	}
	return true, nil
}

func main() {
	nombre, apellido, domicilio := "camila", "gonzalez", "calle 123"
	var dni, telefono int64 = 12345678, 132456

	yaExiste, errExiste := verificarExistencia()
	if errExiste != nil {
		fmt.Printf("error: %s\n", errExiste)
		os.Exit(1)
	}
	if yaExiste == true {
		fmt.Printf("ya existe el cliente\n")
		os.Exit(1)
	}

	_, errValores := valoresValidos(nombre, apellido, domicilio, dni, telefono)
	if errValores != nil {
		fmt.Printf("error: %s\n", errValores)
		os.Exit(1)
	}

	defer fmt.Printf("No han quedado archivos abiertos\n")
	defer fmt.Println("Se detectaron varios errores en tiempo de ejecución")
	defer fmt.Println("Fin de la ejecución")

	_, errLegajo := generarLegajo()
	if errLegajo != nil {
		//panic("no se pudo generar el legajo")
	} else {
		nuevoCliente := Cliente{nombre, apellido, domicilio, 12345, dni, telefono}
		fmt.Printf("Cliente creado con éxito: %v\n", nuevoCliente)
	}

}
