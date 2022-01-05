// 가변인자함수, 복수개의 리턴값

package main

func main() {
	count, total := sum(1, 7, 3, 5, 9)
	println(count, total)
}

func sum(nums ...int) (int, int) {
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}
 
