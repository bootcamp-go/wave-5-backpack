package main

import "fmt"
import "os"
import "strings"
import "strconv"


func main(){
	
	data,_:=os.ReadFile("./salida.csv")
	salida:= "ID\t\tPrecio\tCantidad\n"
	var total=0.0
	for _,linea :=range strings.Split(string(data),"\n") {
		items:=strings.Split(linea,",")
		if(len(items)==3){
			salida=salida+fmt.Sprint(items[0],"\t\t",items[1],"\t",items[2],"\n")
			precio,_:=strconv.ParseFloat(items[1],64)
			cantidad,_:=strconv.ParseFloat(items[2],64)
			total=total+(precio*cantidad)
		}
		
	}
	salida=fmt.Sprintf(salida+"\t\tTotal %f",total)
	fmt.Println(salida)
}

/*
Ejercicio 2 - Leer archivo
La misma empresa necesita leer el archivo almacenado, para ello requiere que: 
se imprima por pantalla mostrando los valores tabulados, con un 
título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), 
el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

Ejemplo:

ID                            Precio  Cantidad
111223                      30012.00         1
444321                    1000000.00         4
434321                         50.50         1
                          4030062.50

*/