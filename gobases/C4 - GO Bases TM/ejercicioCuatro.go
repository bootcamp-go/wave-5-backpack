package main

import(
	"fmt"
	"errors"
)

//Calculo los valores de salario de acuerdo a las horas
//trabajadas y su valor
func calculoSalario(horas int, valorHora float64)(float64,error){
	//Utilizando punteros
	salario := float64(horas) * valorHora
	var salarioMensual *float64

	//Valido condiciones de salario
	switch salarioMensual = &salario;{
	case *salarioMensual > 150000:
		*salarioMensual -= (*salarioMensual * 0.10)
	case *salarioMensual < 80 || horas < 0:
		err := fmt.Errorf("\nerror: ingresó %v horas, el trabajador no puede haber trabajado menos de 80 hs mensuales. \nAdicionalmente, revise que las horas no sean negativas. Horas ingresadas: %v",*salarioMensual, horas)
		return *salarioMensual,err
	}
	
	return *salarioMensual,nil
}

//Para saber cuál de los salarios es el mejor
//Adicionalmente para revisar en caso de valores negativos
func mejorSalario(salarios []float64) (float64,error){
	valueMax := salarios[0]
	for _, valor := range salarios{
		if valor >= 0 {
			if valueMax <= valor {
				valueMax = valor
			}
		}else{
			err := errors.New("error: no pueden haber salarios negativos")
			return valueMax, err
		}
	}
	return valueMax, nil
}

//utilizo la fórmula con el mejor salario obtenido
func calculoAguinaldo(meses int, salarios ...float64)(float64,error){
	value,err := mejorSalario(salarios)
	aguinaldo := 0.0
	if(err != nil){
		return aguinaldo,err
	}else{
		aguinaldo = value / float64(12 * meses)
	}
	return aguinaldo,nil
}

func main(){
	//Cálculo salario
	calculo, err := calculoSalario(90,5000.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Salario: ",calculo)

	//Cálculo aguinaldo
	aguinaldo, errA := calculoAguinaldo(3, 2000000, 3500000, 3200000, 400000, 100000, 2700000)
	if errA != nil {
		fmt.Println(errA)
		return
	}
	fmt.Printf("Aguinaldo: %.0f\n",aguinaldo)
}