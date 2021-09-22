package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1} // Y:0 is implicit
	v3 = Vertex{}     // X:0 and Y:0
	p  = &Vertex{1, 2}
)

func main() {
	// p := &v1
	// fmt.Println(*p)
	// p.X = 4
	fmt.Println(v1, p, v2, v3)
	fmt.Println(p == &v1) // false

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
