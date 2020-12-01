package main

import "fmt"

func main() {
	f := func(x int) {
		fmt.Println("my lucky number is:", x)
	}
	f(13)
}
