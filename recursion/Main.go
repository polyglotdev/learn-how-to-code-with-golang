package main

import "fmt"

func main() {
	fmt.Println("Hello!")
	f := factorial(4)
	fmt.Println(f)
}

func factorial(a int) int {
	if a == 0 {
		return 1
	}
	return a * factorial(a-1)
}
