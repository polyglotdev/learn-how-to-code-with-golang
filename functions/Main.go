package main

import "fmt"

func main() {
	foo()
	bar("Ezra")
	s1 := woo("Hallan")
	fmt.Println(s1)
	x, y := mouse("Elijah", "Hallan")
	fmt.Println(x)
	fmt.Println(y)
	execute(2, 3, 4, 5, 6, 7, 8)
}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := false

	return a, b
}

func execute(x ...int) {
	sum := 0
	for _, v := range x {
		sum += v
		fmt.Println(sum)
	}
}
