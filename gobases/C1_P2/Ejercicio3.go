package main

import "fmt"

func main() {
	var num int = 2
	months := map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril"}
	if num > 0 {
		fmt.Println(num, ":", months[num])
	}
}
