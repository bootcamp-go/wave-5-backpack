package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

type product struct {
	id       int
	price    float64
	quantity int
}

func ReadFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	return data, err
}

func ParseFile(fileData []byte) []product {
	var productList []product
	var productId int
	var productPrice float64
	var productQuantity int
	var err error
	productStringList := strings.Split(string(fileData), "\n")
	for _, productString := range productStringList {
		if productString == "" {
			break
		}
		splitedProductString := strings.Split(productString, ",")
		productId, err = strconv.Atoi(splitedProductString[0])
		if err != nil {
			panic(err.Error())
		}
		productPrice, err = strconv.ParseFloat(splitedProductString[1], 64)
		if err != nil {
			panic(err.Error())
		}
		productQuantity, err = strconv.Atoi(splitedProductString[2])
		if err != nil {
			panic(err.Error())
		}
		productList = append(productList, product{
			id:       productId,
			price:    productPrice,
			quantity: productQuantity,
		})
	}
	return productList
}

func PrintFile(productList []product) {
	w := tabwriter.NewWriter(os.Stdout, 1, 8, 1, '\t', tabwriter.Escape)
	fmt.Fprint(w, "id \t price \t quantity\n")
	var totalPrice float64
	for _, product := range productList {
		totalPrice += product.price * float64(product.quantity)
		fmt.Fprintf(w, " %d\t %.2f \t %d\n", product.id, product.price, product.quantity)
	}
	fmt.Fprintf(w, "\t %.2f \n", totalPrice)
	w.Flush()
}

func main() {
	data, err := ReadFile("products.csv")
	if err != nil {
		fmt.Println("Error when printing data")
	}
	productList := ParseFile(data)
	PrintFile(productList)
}
