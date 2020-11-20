package main

import "fmt"

const (
	a  = iota
	bb = iota
	c  = iota
)

func Iota() {
	fmt.Println(a)
	fmt.Println(bb)
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	xxx := 4
	fmt.Printf("%d\t\t%b", xxx, xxx)

	yyy := xxx << 1
	fmt.Printf("%d\t\t%b", yyy, yyy)
}
