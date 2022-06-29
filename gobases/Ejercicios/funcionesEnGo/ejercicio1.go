package main

func impuestosDeSalario(salario float32) float32 {
	//Ejercicio 1
	var impuesto float32

	if salario > 50000 && salario < 150000 {
		impuesto = (salario * 17) / 100
	} else if salario > 150000 {
		impuesto = (salario * 10) / 100
	} else {
		impuesto = 0
	}

	return impuesto
}
