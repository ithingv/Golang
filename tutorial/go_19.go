// 배열
package main

func main() {
	a := []int{1, 2, 3}
	for s := range a {
		println(s)
	}

	// 정수형 3개 요소를 갖는 배열 a 선언
	var b [3]int
	b[0] = 1
	b[1] = 2
	b[2] = 3
	println(b[1])

	var a1 = [3]int{1, 2, 3}   // 배열의 크기 지정
	var a2 = [...]int{1, 2, 3} // 배열의 크기 자동으로

	for s := range a1 {
		println(s)
	}
	for s := range a2 {
		println(s)
	}

	// 다차원 배열
	var multiArray [3][4][5]int // 정의
	multiArray[0][1][2] = 10    // 사용

	// 다차원 배열의 초기화
	multiArr := func() {
		var a = [2][3]int{
			{1, 2, 3},
			{4, 5, 6}, // 끝에 콤마 추가
		}
		println(a[1][2])
	}
	multiArr()
}
