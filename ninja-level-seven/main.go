package main

import "fmt"

func main() {
	a := 12
	fmt.Println(&a)
	p1 := person{
		firstName: "Dom",
		lastName:  "Hallan",
		age:       37,
	}
	changeMe(&p1)
	fmt.Println(p1)
}

type person struct {
	firstName string
	lastName  string
	age       int
}

func changeMe(p *person) {
	(*p).firstName = "Elijah"
}
