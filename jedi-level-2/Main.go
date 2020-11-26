package main

import (
	"bytes"
	"fmt"
)

const valueone = 356
const valuetwo = 358
const valuethree = 360
const valuefour = -99
const valuefive = -32
const valuesix = -360

var (
	exampleone   = 1 == 86594
	exampletwo   = 2 <= 49
	examplethree = 3 >= 9
	examplefour  = 4 != 10
	examplefive  = 5 > 983
	examplesix   = 6 > 5
)

func main() {
	a := 42
	fmt.Printf("%d\n%b\n%#x\n", a, a, a)

	fmt.Println(exampleone)
	fmt.Println(exampletwo)
	fmt.Println(examplethree)
	fmt.Println(examplefour)
	fmt.Println(examplefive)
	fmt.Println(examplesix)

	fmt.Println(valueone)
	fmt.Println(valuetwo)
	fmt.Println(valuethree)
	fmt.Println(valuefour)
	fmt.Println(valuefive)
	fmt.Println(valuesix)

	rs := bytes.Runes([]byte("Dominique"))
	for _, r := range rs {
		fmt.Printf("%#U\n", r)
	}

	d := 42
	fmt.Printf("%d\t%b\t%U\n", d, d, d)

	d = d << 1
	fmt.Println(d)
}
