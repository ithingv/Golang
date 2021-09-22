package main

import "fmt"

var (
	a = []int{1, 2, 3, 4, 5}
)

func main() {
	// array
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// slice
	b := a[1:4]
	for i := 0; i < len(b); i++ {
		b[i] = 100
		fmt.Println(a[i+1]) // Slices are like references to arrays
	}
	fmt.Println(len(b), cap(b))
}
