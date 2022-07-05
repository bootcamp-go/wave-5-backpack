package main

import "fmt"

type ListaHeterogenea struct {
	Len int
	Data []interface{}
}

func main() {
	
	l := ListaHeterogenea{}
	l.Data = append(l.Data, 1)
	l.Data = append(l.Data, "hola")
	l.Data = append(l.Data, true)

	fmt.Printf("%v\n", l.Data)

}
