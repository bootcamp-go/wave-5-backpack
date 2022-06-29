package main

func calcularSalario(mPorMes float32, categoria string) float32 {
	//Ejercicio 3
	hPorMes := mPorMes / 60
	var salario float32

	switch categoria {
	case "C":
		salario = hPorMes * 1000
	case "B":
		salario = (hPorMes * 1500) + (hPorMes*1500)*20/100
	case "A":
		salario = (hPorMes * 3000) + (hPorMes*3000)*50/100
	}

	return salario
}
