package main

import "fmt"

func Strings() {
	fmt.Printf("%T\n", "Hello")
	bs := []byte("Dom")
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	s := "Dominique Israel Hallan"

	for i := 0; i < len(bs); i++ {
		fmt.Printf("%#U\n", bs[i])
	}

	for i, v := range s {
		fmt.Printf("at index position %d we have %d\n", i, v)
	}
}
