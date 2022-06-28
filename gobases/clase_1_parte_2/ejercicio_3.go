package main

import "fmt"

var meses = map[int]string{1: "ENERO", 2: "FEBRERO", 3: "MARZO", 4: "ABRIL", 5:"MAYO", 6:"JUNIO", 7:"JULIO", 8:"AGOSTO", 9:"SEPTIEMBRE", 10:"OCTUBRE", 11:"NOVIEMBRE", 12:"DICIEMBRE"}

func main(){
	var rec int = 4
	fmt.Println(meses[rec])
	
}