package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 46, 45, 40, 41, 42}
	fmt.Println(arr)

	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("type: %T\n", arr)

	m := map[string]string{
		"Name":         "Dom",
		"Occupation":   "Software Engineer in training",
		"Hates Coding": "Yes",
	}

	delete(m, "Hates Coding")
	fmt.Println(m)
}
