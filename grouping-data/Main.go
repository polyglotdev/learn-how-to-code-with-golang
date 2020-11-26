package main

import (
	"fmt"
)

func main() {
	var x [5]int
	fmt.Println(x)

	y := []int{5, 9, 14, 13, 69}
	fmt.Println(y[1]) // this should give back 9
	fmt.Println(y[1:3])
	fmt.Println(len(y))

	for i, v := range y {
		fmt.Printf("value: %v at index: %v\n", v, i)
	}

	// using a for loop access every indices value without using range
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
	y = append(y, 1, 2, 3, 4, 5, 699)
	fmt.Println(y)

	z := []int{234, 345, 333, 973}
	y = append(y, z...)
	fmt.Println(y)

	y = append(y[:2], y[4:]...)
	fmt.Println(y)

	zz := make([]int, 10, 100)
	fmt.Println(zz)
	fmt.Println(len(zz))
	fmt.Println(cap(zz))

	zz = append(zz, 4594)
	fmt.Println(zz)

	eh := []string{"Elijah", "Hallan", "Chocolate"}
	ezh := []string{"Ezra", "Hallan", "Chocolate"}

	xp := [][]string{eh, ezh}
	fmt.Println(xp)

	phonenumber := map[string]int{
		"Becky":   3147483948,
		"Kristen": 6187364536,
	}

	fmt.Println(phonenumber)
	fmt.Println(phonenumber["Becky"])

	if v, ok := phonenumber["Becky"]; ok {
		fmt.Println("This is is the IF print", v)
	}

	phonenumber["Dom"] = 3142238874

	for i, v := range phonenumber {
		fmt.Println(i, v)
	}

	delete(phonenumber, "Becky")
	fmt.Println(phonenumber)
}
