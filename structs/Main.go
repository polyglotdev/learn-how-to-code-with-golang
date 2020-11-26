package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Dom",
		last:  "Hallan",
		age:   1,
	}

	fmt.Println(p1.first, p1.last, p1.age)
}
