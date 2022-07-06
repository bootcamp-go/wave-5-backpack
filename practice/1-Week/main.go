package main

import "fmt"

type Xpoison interface {
	Walk() string
}
type Ypoison interface {
	Swin() string
}
type Zpoison interface {
	Fly() string
}

type Homunculus struct {
	Name string
}

func (this Homunculus) Walk() string {
	x := fmt.Sprintf("%s, now can Walk", this.Name)
	return x
}
func (this Homunculus) Swim() string {
	x := fmt.Sprintf("%s, now can Swim", this.Name)
	return x
}

func main() {
	EdwardElric := Homunculus{Name: "Edward Elric"}
	fmt.Println(EdwardElric.Swim())
}
