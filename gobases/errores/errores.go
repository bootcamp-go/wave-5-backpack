package main

import (
	"fmt"
)

type miError struct {
	mensaje string
}

func (e *miError) Error() string {
	return e.mensaje
}

func main() {

	//Ejercicio 1
	// En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
	// Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario ingresado no alcanza el mínimo imponible"
	// y lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.

	fmt.Println("++++++++ EJERCICIO 1,2,3 ++++++++++")
	fmt.Println("INGRESE UN SALARIO:")

	var salary float64
	var punteroASalary *float64
	punteroASalary = &salary

	fmt.Scanf("%f", &salary)

	//DIRECCION DE MEMORIA DE SALARY
	fmt.Println("&", &salary)
	//PUNTERO A SALARY QUE RETORNA EL VALOR EN LA DIRECCION DE MEMORIA
	fmt.Println("*", *punteroASalary)

	respuesta, err := CalcularImpuesto(punteroASalary)

	if err != nil {
		fmt.Println(err.Error())

	} else {
		fmt.Println(respuesta)
	}

	//Ejercicio 4

	fmt.Println("++++++++ EJERCICIO 4 ++++++++++")
	fmt.Println("INGRESE LAS HORAS TRABAJADAS:")

	var horasTrabajadas int
	var punteroAHorasTrabajadas *int
	punteroAHorasTrabajadas = &horasTrabajadas

	fmt.Scanf("%d", &horasTrabajadas)

	fmt.Println("INGRESE EL VALOR DE LAS HORAS:")

	var valorHora float64
	var punteroAValorHora *float64
	punteroAValorHora = &valorHora

	fmt.Scanf("%f", &valorHora)

	sueldo, err := CalcularSueldo(punteroAHorasTrabajadas, punteroAValorHora)

	if err != nil {
		fmt.Println(err.Error())
	}

	var punteroASueldo *float64 = &sueldo
	aguinaldo, erro := CalcularAguinaldo(*punteroASueldo)

	if erro != nil {
		fmt.Println(erro.Error())
	} else {
		fmt.Println("El aguinaldo es:", aguinaldo)
	}

}

func CalcularImpuesto(salary *float64) (string, error) {
	if *salary < 150000 {
		return "", fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %.2f.Fecha", *salary) //errors.New("error: el salario ingresado no alcanza el mínimo imponible") //&miError{mensaje: "error: el salario ingresado no alcanza el mínimo imponible"}
	} else {
		return "Debe pagar impuesto", nil
	}
}

//Ejercicio 4
func CalcularSueldo(horasTrabajadas *int, valorHora *float64) (float64, error) {

	var sueldo float64 = float64(*horasTrabajadas) * *valorHora

	if sueldo >= 150000 {
		sueldoAux := sueldo
		sueldo = sueldo - (sueldo * .10)
		fmt.Printf("Al sueldo de %.2f se le desconto un .10 quedando: %.2f\n", sueldoAux, sueldo)
		return sueldo, nil
	}

	if *horasTrabajadas < 80 {
		return 0, fmt.Errorf("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	return 0, nil
}

func CalcularAguinaldo(sueldo float64) (float64, error) {

	if sueldo < 0 {
		return 0, fmt.Errorf("error: el salario ingresado no puede ser negativo")
	} else {
		return (sueldo / 12) * 6, nil
	}

}
