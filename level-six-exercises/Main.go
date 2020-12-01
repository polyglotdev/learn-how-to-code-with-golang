package main

import "fmt"

// func main() {
// 	a := foo()
// 	b, c := bar()

// 	fmt.Println(a)
// 	fmt.Println(b, c)
// }

// func foo() int {
// 	return 4
// }

// func bar() (int, string) {
// 	return 4, "Hello"
// }

// func main() {
// 	fmt.Println("Hands On #2")
// 	aa := []int{1, 2, 3, 4, 5}
// 	a := add(aa...)
// 	fmt.Println(a)

// 	bb := []int{1, 2, 3, 4, 5}
// 	b := bar(bb)
// 	fmt.Println(b)
// }

// func add(a ...int) int {
// 	total := 0
// 	for _, v := range a {
// 		total += v
// 	}
// 	return total
// }

// func bar(xi []int) int {
// 	total := 0
// 	for _, v := range xi {
// 		total += v
// 	}
// 	return total
// }

// func main() {
// 	defer foo()
// 	bar()
// }

// func foo() {
// 	fmt.Println("#1 stunna")
// }

// func bar() {
// 	fmt.Println("#2 stunna -- betcha I print tho")
// }

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last, "and I am", p.age, "years old.")
}

func main() {
	fmt.Println("Hands on #4")

	p1 := person{
		first: "Ezra",
		last:  "Hallan",
		age:   3,
	}
	fmt.Println(p1)
	p1.speak()
}
