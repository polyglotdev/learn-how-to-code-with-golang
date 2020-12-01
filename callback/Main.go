package main

import "fmt"

func main() {
	fmt.Println("Hello, playground")
	sumResult := sum(1, 2, 3, 4, 5)
	fmt.Println(sumResult)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
