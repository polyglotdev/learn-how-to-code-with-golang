package main

import "fmt"

func main() {
	var x int
	fmt.Println(x)

	x++
	fmt.Println(x)

	foo()

	incrementor()

}

func foo() {
	fmt.Println("hello")
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		fmt.Println(x)
		return x
	}
}
