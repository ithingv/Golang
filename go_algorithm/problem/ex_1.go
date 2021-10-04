package main

import (
	"fmt"
)

func gcd(n, m int) int {
	for m != 0 {
		n, m = m, n&m
	}
	return n
}

func solution(n int, m int) []int {
	return []int{}
}

func main() {
	fmt.Println(gcd(3, 12))
}
