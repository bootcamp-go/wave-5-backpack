package main

import (
	"fmt"
	"os"
)

func main() {
  data := fmt.Sprintf("1; 35.5; 5")

  os.WriteFile("producto.csv", []byte(data), 0644)
}
