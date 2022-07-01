package main

import "fmt"

type user struct{
	name string
	lastName string
	age int
	email string
	password string
}

func main(){

	person :=user{"nicolas","rojas",25,"nicolas@ottis.com","moteros2022"}
	
	schangeEmail(person,"nicoemoxito22@gmail.com")

	fmt.Println(person)
}

func changeName(p *user,name string,lastName string){
	p.name=name	
	p.lastName=lastName
}


func changeAge(p *user,age string){
	p.age=age
}
func changeEmail(p *user,mail string){
	p.email=mail
}
func changePassword(p *user,password string){
	p.password=password
}