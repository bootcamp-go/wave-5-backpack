package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type producto struct {
	Id       int
	precio   int
	cantidad int
}

func int_2_str(num int) string {
	return fmt.Sprintf("%d", num)
}

func str_2_int(pal_num string) int {
	num, e := strconv.Atoi(pal_num)
	if e == nil {
		fmt.Printf("%T \n %v", num, num)
	}
	return num
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}
func readFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func write_string(texto, direccion string) {
	fmt.Println(texto)
	data := []byte(texto)
	fmt.Println(data)
	err := os.WriteFile(direccion, data, 0644)
	fmt.Println(err)
}
func main() {
	fmt.Println("hoal")
	// lista_super_mercado := [][]int{{0, 14, 3}, {1, 5, 3}}
	// list_2_string := "ID;precio;cantidad\n"
	// for _, list := range lista_super_mercado {
	// 	list_2_string += arrayToString(list, ";") + "\n"
	// }
	// write_string(list_2_string, "lista_super_mercado.csv")
	csvFile, err := os.Open("lista_super_mercado.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
	csvLines := csv.NewReader(csvFile) //.ReadAll()
	csvLines.Comma = ';'
	reader, _ := csvLines.ReadAll()
	fmt.Println("printiar lineas")
	list_prod := []producto{}
	indices := []string{}
	fmt.Println(indices)
	for pos, line := range reader {
		if pos == 0 {
			indices = line
			continue
		}
		prod := producto{str_2_int(line[0]), str_2_int(line[1]), str_2_int(line[2])}
		list_prod = append(list_prod, prod)
		fmt.Println(line[0])
		fmt.Println(line)
	}
	fmt.Println(list_prod[1].precio)
	total := 0
	for _, elem := range list_prod {
		total += elem.precio
	}
	matriz := form_table(indices, list_prod, total)
	fmt.Println(matriz)
	imprimir_prod(matriz)
}

func form_table(indices []string, list_prod []producto, total int) [][]string {
	matriz := [][]string{}
	matriz = append(matriz, indices)
	for _, elem := range list_prod {
		matriz = append(matriz, []string{int_2_str(elem.Id), int_2_str(elem.precio), int_2_str(elem.cantidad)})
	}
	matriz = append(matriz, []string{"", int_2_str(total), ""})
	return matriz
}

func imprimir_prod(table [][]string) {
	for pos, elem := range table {
		if pos == 0 {
			fmt.Printf("%s\t%s %s\n", elem[0], elem[1], elem[2])
			continue
		}
		fmt.Printf("%s\t    %s\t      %s\n", elem[0], elem[1], elem[2])
	}
}
