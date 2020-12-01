package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	name string
	age  int
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I called secret agent - ", s.name)
}

func (p person) speak() {
	fmt.Println("I am", p.name)
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

func main() {
	p1 := secretAgent{
		person: person{
			"Dom Hallan",
			37,
		},
		ltk: true,
	}

	p2 := person{
		name: "Elijah Hallan",
		age:  5,
	}

	fmt.Println(p1.name)
	p1.speak()

	fmt.Println(p2)
	p2.speak()

	bar(p1)
	bar(p2)
}
