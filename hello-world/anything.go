package main

import "fmt"

func main() {
	count := 100
	fmt.Println("Hello, World!")

	for i := 0; i <= count; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

/*
  control flow:
  1. sequence
  2. loop; iterative
  3. conditional
*/
