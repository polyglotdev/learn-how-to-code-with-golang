package main

import "fmt"

type person struct {
	name          string
	age           int
	occupation    string
	martialStatus bool
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.name)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James Bond",
			41,
			"Spy",
			false,
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Dom Hallan",
			37,
			"Software Engineer",
			true,
		},
		ltk: false,
	}
	fmt.Println(sa1)
	sa1.speak()

	fmt.Println(sa2)
	sa2.speak()
}
