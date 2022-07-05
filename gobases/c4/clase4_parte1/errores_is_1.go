package main

import (
	"errors"
	"fmt"
)

var err1 = errors.New("error número 1")

func myFun() error {
	return fmt.Errorf("información extra del error: %w", err1)
}

func main() {
	e := myFun()
	coincidence := errors.Is(e, err1)
	fmt.Println(coincidence)
}
