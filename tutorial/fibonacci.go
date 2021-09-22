package main

import "fmt"

func fibonacci1() func() int {
	a, b := 1, 0
	return func() int {
		a, b = b, a+b
		return a
	}
}

func fibonacci2() func() int {
	a, b := 0, 1
	return func() int {
		f := a
		a, b = b, b+f
		return f
	}
}

func main() {
	f := fibonacci1() // fibonacci2()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
