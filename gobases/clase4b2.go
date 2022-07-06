package main

import (
	"errors"
	"fmt"
	"os"
)

type cliente struct {
	legado          int
	nombre_completo string
	DNI             string
	numTelefono     int
	domicilio       string
}

var contadorId int

func generadorLegado(aleatorio int) *int {
	if aleatorio == 1 {
		contadorId += 1
		return &contadorId
	}
	return nil
}
func finalGeneradorLegado() int {
	numero := generadorLegado(1)
	if numero == nil {
		panic("todo salio mal aborten")
	}
	return (*numero)
}
func readFile(path string) string {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("el archivo estaba corrupto")
			return
		}
	}()
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}
func verificadorDeNulos(c cliente) (*cliente, error) {
	if c.DNI == "" {
		return nil, errors.New("error en el dni nulo")
	}
	if c.domicilio == "" {
		return nil, errors.New("error domicilio nulo")
	}
	if c.legado == 0 {
		return nil, errors.New("error legado cero")
	}
	if c.nombre_completo == "" {
		return nil, errors.New("error nombre nulo")
	}
	if c.numTelefono == 0 {
		return nil, errors.New("error telfono nulo")
	}
	return &c, nil
}

func main() {
	fmt.Println("hoal")
	legado := finalGeneradorLegado()
	fmt.Println(legado)
	textoVerificadoCliente := readFile("./clientes.csv")
	fmt.Println(textoVerificadoCliente)
	c := cliente{
		legado:          legado,
		nombre_completo: "polo",
		DNI:             "A456633",
		numTelefono:     98473839,
		domicilio:       "los membrillos",
	}

	a, err := verificadorDeNulos(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*a)
	defer func() {
		fmt.Println("Fin de la ejecución")
		fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		fmt.Println("No han quedado archivos abiertos")

	}()
	panic("mierda")

}
