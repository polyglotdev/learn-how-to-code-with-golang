package main

import "fmt"

func main() {
	pp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(pp...)
	fmt.Println(s)

	bb := even(sum, pp...)
	fmt.Println(bb)

	cc := odd(sum, pp...)
	fmt.Println(cc)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%3 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
