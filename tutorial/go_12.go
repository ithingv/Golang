// 가변인자함수

package main

func main() {
	total := sum(1, 7, 3, 5, 9)
	println(total)
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}
