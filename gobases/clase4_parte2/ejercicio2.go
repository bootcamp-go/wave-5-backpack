package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Cliente struct {
	Legajo           int
	Nombre, Apellido string
	DNI, Telefono    int
	Domicilio        string
}

func getID() int {
	rand.Seed(time.Now().Unix())
	random := (rand.Float64() - 0.5) * 10000
	id := int(math.Round(random)) // En la mitad de los casos arroja panic
	if id > 0 {
		// Verificamos que el cliente exista
		fileName := "./customers.txt"
		_ = leerArchivo(fileName)
		fmt.Println("seguimos  despues del panic :)")
		return id
	} else {
		panic("No pudo generarse el ID " + strconv.Itoa(id))
	}
}

func leerArchivo(fileName string) string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	data, err := os.ReadFile(fileName)
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	} else {
		return string(data)
	}
}

func (cliente *Cliente) verificarDatos(Nombre, Apellido, Domicilio string, DNI, Telefono int) (string, error) {
	if Nombre == "" || Apellido == "" || Domicilio == "" || DNI < 0 || Telefono < 0 {
		return "", errors.New("Algunos de los datos son nulos")
	} else {
		cliente.Nombre = Nombre
		cliente.Apellido = Apellido
		cliente.Domicilio = Domicilio
		cliente.DNI = DNI
		cliente.Telefono = Telefono
		return "Registro exitoso", nil
	}
}

func main() {
	defer func() {
		fmt.Println("Fin de la ejecucion")
		fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		fmt.Println("No han quedado archivos abiertos")
	}()

	legajo := getID()
	cliente1 := Cliente{Legajo: legajo}

	cliente1.verificarDatos("Camilo", "Calderon", "Colombia", 123, 456)
	fmt.Println(cliente1)
}
