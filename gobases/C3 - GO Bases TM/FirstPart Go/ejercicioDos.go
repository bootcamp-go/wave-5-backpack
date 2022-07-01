package main

import(
	"fmt"
	"strings"
	"strconv"
	"os"
)

func main(){
	//Lectura del archivo csv
	data, _ := os.ReadFile("./infoProducto.csv")
	
	var precio,_ []int
	var cantidad,_ []int
	total, counter := 0, 4

	//Para imprimir los datos del csv de manera tabulada
	//y para guardar los datos correspondientes a cantidad y precio
	for i,info := range strings.Fields(string(data)){
		stringCsv := strings.Split(info, ",")
		if (i+1) % 3 == 0{
			valueCantidad,_ := strconv.Atoi(stringCsv[0])
			cantidad = append(cantidad, valueCantidad)
			fmt.Print(stringCsv[0],"\n")
		}else{
			fmt.Print(stringCsv[0],"\t")
		}
		if i == counter {
			valuePrecio,_ := strconv.Atoi(stringCsv[0])
			precio = append(precio, valuePrecio)
			counter += 3
		}
	}

	//Recorro precio y cantidad para calcular el total
	for k:=0;k<len(precio);k++{
		total += precio[k] * cantidad[k]
	}
	fmt.Println("Total\t",total)
}