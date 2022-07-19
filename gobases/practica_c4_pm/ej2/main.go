package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Cliente struct {
	Legajo         int
	Nombre         string
	Apellido       string
	DNI            int64
	NumeroTelefono int64
	Domicilio      string
}

type ErrorCliente struct {
	nombre   string
	apellido string
	legajo   int
}

func (this *ErrorCliente) Error() string {
	return fmt.Sprintf("No se pudo agregar el cliente: %s %s con legajo %d", this.nombre, this.apellido, this.legajo)
}

func generarID() int {
	return rand.Intn(100)
}

func leerArchivo(archivo string) []Cliente {
	file, err := os.ReadFile(archivo)

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

	filas := strings.Split(string(file), "\n")

	var clientes []Cliente = []Cliente{}
	for _, fila := range filas {
		if fila == "" {
			continue
		}

		datos := strings.Split(fila, ";")

		legajo, _ := strconv.Atoi(datos[0])
		dni, _ := strconv.ParseInt(datos[3], 10, 64)
		numeroTelefono, _ := strconv.ParseInt(datos[4], 10, 64)

		cliente := Cliente{
			Legajo:         legajo,
			Nombre:         datos[1],
			Apellido:       datos[2],
			DNI:            dni,
			NumeroTelefono: numeroTelefono,
			Domicilio:      datos[5],
		}
		clientes = append(clientes, cliente)
	}
	return clientes
}

func verificarExistencia(cliente Cliente) bool {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error: ", err)
		}
	}()

	clientes := leerArchivo("customers.txt")
	for _, c := range clientes {
		if c.DNI == cliente.DNI {
			return true
		}
	}
	return false
}

func validarDatos(cliente Cliente) (bool, error) {

	contErrores := 1
	var err error = nil

	if cliente.Legajo == 0 {
		err = fmt.Errorf("%d. El legajo no puede ser 0", contErrores)
		contErrores++
	}

	if cliente.Nombre == "" {
		if contErrores-1 == 0 {
			err = fmt.Errorf("%d. El nombre no puede ser vacio", contErrores)
		} else {
			err = fmt.Errorf("%d. El nombre no puede ser vacio\n%w", contErrores, err)
		}
		contErrores++
	}

	if cliente.Apellido == "" {
		if contErrores-1 == 0 {
			err = fmt.Errorf("%d. El apellido no puede ser vacio", contErrores)
		} else {
			err = fmt.Errorf("%d. El apellido no puede ser vacio\n%w", contErrores, err)
		}
		contErrores++
	}

	if cliente.DNI == 0 {
		if contErrores-1 == 0 {
			err = fmt.Errorf("%d. El DNI no puede ser 0", contErrores)
		} else {
			err = fmt.Errorf("%d. El DNI no puede ser 0\n%w", contErrores, err)
		}
		contErrores++
	}

	if cliente.NumeroTelefono == 0 {
		if contErrores-1 == 0 {
			err = fmt.Errorf("%d. El numero de telefono no puede ser 0", contErrores)
		} else {
			err = fmt.Errorf("%d. El numero de telefono no puede ser 0\n%w", contErrores, err)
		}
		contErrores++
	}

	if cliente.Domicilio == "" {
		if contErrores-1 == 0 {
			err = fmt.Errorf("%d. El domicilio no puede ser vacio", contErrores)
		} else {
			err = fmt.Errorf("%d. El domicilio no puede ser vacio\n%w", contErrores, err)
		}
		contErrores++
	}

	return contErrores-1 == 0, err

}

func agregarCliente(cliente Cliente) error {
	strCliente := fmt.Sprintf("%d;%s;%s;%d;%d;%s\n", cliente.Legajo, cliente.Nombre, cliente.Apellido, cliente.DNI, cliente.NumeroTelefono, cliente.Domicilio)
	err := os.WriteFile("customers.txt", []byte(strCliente), 0644)
	if err != nil {
		return &ErrorCliente{cliente.Nombre, cliente.Apellido, cliente.Legajo}
	}
	return nil
}

func main() {
	defer func() {
		fmt.Println("Fin de la ejecución")
		fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		fmt.Println("No han quedado archivos abiertos")
	}()

	cliente := Cliente{
		Legajo:         generarID(),
		Nombre:         "Matias",
		Apellido:       "Carrasco",
		DNI:            159600793,
		NumeroTelefono: 951126368,
		Domicilio:      "Rey Felipe II 024, Estacion Central",
	}

	fmt.Println(cliente)

	if verificarExistencia(cliente) {
		fmt.Println("El cliente ya existe")
	} else {
		fmt.Println("El cliente no existe")
		verificacion, err := validarDatos(cliente)
		if err != nil {
			fmt.Println(err)
		}

		if verificacion {
			fmt.Println("El cliente es válido")
			err := agregarCliente(cliente)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El cliente se agregó correctamente")
			}
		} else {
			fmt.Println("El cliente no es válido")
		}
	}

	fmt.Println()
	cliente = Cliente{
		Legajo:         generarID(),
		Nombre:         "",
		Apellido:       "",
		DNI:            0,
		NumeroTelefono: 0,
		Domicilio:      "",
	}

	fmt.Println(cliente)

	if verificarExistencia(cliente) {
		fmt.Println("El cliente ya existe")
	} else {
		fmt.Println("El cliente no existe")
		verificacion, err := validarDatos(cliente)
		if err != nil {
			fmt.Println(err)
		}

		if verificacion {
			fmt.Println("El cliente es válido")
			err := agregarCliente(cliente)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El cliente se agregó correctamente")
			}

		} else {
			fmt.Println("El cliente no es válido")
		}
	}

}
