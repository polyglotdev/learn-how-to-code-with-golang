package main

import "fmt"

func main() {
	fmt.Println("Functions")
	Add(10, 10)
}

func Add(a, c int) {
	fmt.Println(a + c)
}
