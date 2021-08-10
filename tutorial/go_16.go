// 일급함수

package main

func main() {
	// 변수 add에 익명 할당
	add := func(i int, j int) int {
		return i + j
	}
	// add 함수 전달
	r1 := add(1, 2)
	println(r1)

	// 직접 첫번째 파라미터에 익명함수를 정의
	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}
