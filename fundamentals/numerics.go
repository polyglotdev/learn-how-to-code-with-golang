package main

import "fmt"

var xx int16
var yy uint32

func Numerics() {
	fmt.Println("Numeric types")
	fmt.Println("--------------")
	fmt.Println(4)

	xx = 32767
	fmt.Printf("%T\n", xx)

	yy = 42929672
	fmt.Printf("%T\n", yy)
}
