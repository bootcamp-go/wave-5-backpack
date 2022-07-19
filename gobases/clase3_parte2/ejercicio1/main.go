package main

import "fmt"

type Usser struct {
	name   string
	l_name string
	age    int
	email  string
	pass   string
}

func (u *Usser) setPrimaryValues(name, l_name, email, pass *string, age *int) {
	u.name = *name
	u.l_name = *l_name
	u.age = *age
	u.email = *email
	u.pass = *pass
}

func (u *Usser) updateName(name, l_name *string) {
	u.name = *name
	u.l_name = *l_name
}
func (u *Usser) updateAge(age *int) {
	u.age = *age
}
func (u *Usser) updateEmail(email *string) {
	u.email = *email
}
func (u *Usser) updatePass(pass *string) {
	u.pass = *pass
}
func (u *Usser) print() {
	fmt.Printf("%s %s %d %s %s\n", u.name, u.l_name, u.age, u.email, u.pass)
}

func main() {
	u := Usser{}
	name1 := "Michael"
	l_name1 := "Torres"
	age1 := 23
	email1 := "michaelstiven.torres@mercadolibre.com.co"
	pass1 := "12345"

	name2 := "Camila"
	l_name2 := "Alvarado"
	age2 := 21
	email2 := "cami.alvarado@mercadolibre.com.co"
	pass2 := "54321"

	u.setPrimaryValues(&name1, &l_name1, &email1, &pass1, &age1)

	u.print()

	u.updateName(&name2, &l_name2)
	u.updateAge(&age2)
	u.updateEmail(&email2)
	u.updatePass(&pass2)

	u.print()
}
