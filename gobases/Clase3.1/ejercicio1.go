package main

import "os"

func main() {
	s := "id;price;quantity\n"
	s += "30;1000;3\n"
	s += "90;5040;20\n"
	s += "10;400;1\n"
	s += "3;99999;2\n"
	s += "7;4500;300\n"

	os.WriteFile("products.csv", []byte(s), 0644)
}
