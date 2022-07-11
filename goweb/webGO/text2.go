package main

import (
	"fmt"
	"reflect"
)

type hola struct {
	cara    int
	pelo    int
	dientes string
}

func main() {

	atrib := map[string]int{
		"cara":    0,
		"pelo":    1,
		"dientes": 20,
	}
	fmt.Println("hola")
	h := hola{2, 3, "nada"}
	// e := reflect.ValueOf(h)
	t := reflect.TypeOf(h)
	fields := reflect.VisibleFields(reflect.TypeOf(struct{ hola }{}))
	for _, field := range fields {
		fmt.Printf("Key: %s\tType: %s\n", field.Name, field.Type)
	}
	strNumFields := t.NumField()
	for i := 0; i < strNumFields; i++ {
		field := t.Field(i)

		fmt.Printf("Field Type: %s: %s Kind: %s\n", field.Name, field.Type.Name(), field.Type.Kind())
	}
	// fmt.Println(e)
	// fmt.Println(t)
	firstNames := make([]string, 0, len(atrib))
	for j := range atrib {
		firstNames = append(firstNames, j)
	}
	fmt.Println("polo")
	for _, k := range firstNames {
		fmt.Println(atrib[k])
	}
	fmt.Println("revisar")

	// t := reflect.TypeOf(h)
	// strNumFields := t.NumField()
	// v := reflect.ValueOf(h)
	// for i := 0; i < v.NumField(); i++ {
	// 	fmt.Println(v.Type().Field(i).Name)
	// 	fmt.Println("\t", v.Field(i))
	// 	fmt.Println(t.Field(i).Type.Kind())
	// 	if t.Field(i).Type.Kind().String() == "int" {
	// 		fmt.Println("hola23232")
	// 	}
	// 	if v.Field(i).String() == "nada" {
	// 		fmt.Println("hoa")
	// 	}
	// }

}
