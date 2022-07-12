package main

import (
	"fmt"
	"os"
)
type Producto struct {
	id int
	precio float64
	cantidad int
}
func main(){
	p1 := Producto{1,3.5,2}
	p2 := Producto{2,4,1}
	p3 := Producto{3,1.5,3}

	productos :=  []Producto{p1,p2,p3}

	txt := ParsString(productos)
	bites := []byte(txt)
	err := os.WriteFile("./productos.csv",bites,0644) 
	if err != nil {
		fmt.Print(err)
	}
}

func ParsString(prod []Producto)string {
	txt := fmt.Sprintf("ID;PRECIO;CANTIDAD\n")
	for _,val := range prod {
		txt += fmt.Sprintf("%d;%.2f;%d\n",val.id,val.precio,val.cantidad)
	}
	return txt
}