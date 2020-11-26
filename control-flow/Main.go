package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("outer loop:%d\t", i)
		for j := 0; j <= 3; j++ {
			fmt.Printf("The inner loop: %d\n", j)
		}
	}

	a := 20
	b := 42

	for a <= b {
		a *= 20
		fmt.Println(a)
	}
}
