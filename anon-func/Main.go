package main

import "fmt"

func main() {
	foo()
	fmt.Println("Hello, playground")

	func(x int) {
		fmt.Println("the meaning of life:", x)
	}(42)
}

func foo() {
	fmt.Println("foo ran")
}
